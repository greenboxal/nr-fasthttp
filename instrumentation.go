package nrfasthttp

import (
	"time"

	"github.com/newrelic/go-agent"
	fasthttp "gopkg.in/valyala/fasthttp.v20160316"
)

// UserValue name used to store newrelic.Transaction on fasthttp.RequestCtx
const NewRelicTransaction = "__newrelic_transaction__"

func StartTransaction(app newrelic.Application, name string, ctx *fasthttp.RequestCtx) newrelic.Transaction {
	return app.StartTransaction(name, nil, &Request{&ctx.Request})
}

func GetTransaction(ctx *fasthttp.RequestCtx) newrelic.Transaction {
	val := ctx.UserValue(NewRelicTransaction)

	if val == nil {
		return nil
	}

	return val.(newrelic.Transaction)
}

func WrapHandler(app newrelic.Application, name string, handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		txn := StartTransaction(app, name, ctx)

		defer func() {
			err := recover()

			if err != nil {
				switch err := err.(type) {
				case error:
					txn.NoticeError(err)
				default:
					txn.NoticeError(errWrapper{err})
				}
			} else {
				txn.ResponseSent(NewResponse(&ctx.Response, &ctx.Request))
			}

			txn.End()

			if err != nil {
				panic(err)
			}
		}()

		ctx.SetUserValue(NewRelicTransaction, txn)

		handler(ctx)
	}
}

func StartExternalSegment(txn newrelic.Transaction, req *fasthttp.Request) newrelic.ExternalSegment {
	return newrelic.StartExternalSegment(txn, Request{req})
}

func Do(client Client, txn newrelic.Transaction, req *fasthttp.Request, res *fasthttp.Response) error {
	seg := StartExternalSegment(txn, req)
	defer seg.End()

	seg.Response = NewResponse(res, req)
	err := client.Do(req, res)

	return err
}

func DoTimeout(client Client, txn newrelic.Transaction, req *fasthttp.Request, res *fasthttp.Response, timeout time.Duration) error {
	seg := StartExternalSegment(txn, req)
	defer seg.End()

	seg.Response = NewResponse(res, req)
	err := client.DoTimeout(req, res, timeout)

	return err
}

func DoDeadline(client Client, txn newrelic.Transaction, req *fasthttp.Request, res *fasthttp.Response, deadline time.Time) error {
	seg := StartExternalSegment(txn, req)
	defer seg.End()

	seg.Response = NewResponse(res, req)
	err := client.DoDeadline(req, res, deadline)

	return err
}
