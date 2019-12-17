// Copyright 2018 NetApp, Inc. All Rights Reserved.

package rest

import (
	"net/http"
	"time"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func Logger(inner http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestId := xid.New()
		logRestCallInfo("REST API call received.", r, start, requestId, routeName)

		inner.ServeHTTP(w, r)

		restOpsTotal.WithLabelValues(r.Method, routeName).Inc()
		endTime := float64(time.Since(start).Milliseconds())
		restOpsSecondsTotal.WithLabelValues(r.Method, routeName).Observe(endTime)

		logRestCallInfo("REST API call complete.", r, start, requestId, routeName)
	})
}

func logRestCallInfo(msg string, r *http.Request, start time.Time, requestId xid.ID, name string) {
	log.WithFields(log.Fields{
		"requestID": requestId,
		"method":    r.Method,
		"uri":       r.RequestURI,
		"route":     name,
		"duration":  time.Since(start),
	}).Debug(msg)
}
