package main

import (
	"flag"
	"log"

	"consumers/slack/utils"

	"github.com/golang/protobuf/ptypes"
	"github.com/thought-machine/dracon/consumers"
)

var (
	Webhook    string // the webhook url to post to
	LongFormat bool   //boolean, False by default, if set to True it dumps all findings in JSON format to the webhook url
)

func main() {
	flag.StringVar(&Webhook, "Webhook", "", "the Webhook to push results to")
	flag.BoolVar(&LongFormat, "long", true, "post the full results to Webhook, not just metrics")

	if err := consumers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	if Webhook == "" {
		log.Fatal("Webhook is undefined")
	}
	if consumers.Raw {
		responses, err := consumers.LoadToolResponse()
		if err == nil {
			log.Fatal(err)
		}
		if !LongFormat {
			messages, err := utils.ProcessRawMessages(responses)
			if err == nil {
				log.Fatal(err)
			}
			for _, msg := range messages {
				utils.PushMessage(msg, Webhook)
			}
		} else {
			scanInfo := utils.GetRawScanInfo(responses[0])
			msgNo := utils.CountRawMessages(responses)
			if tstamp, err := ptypes.Timestamp(scanInfo.GetScanStartTime()); err != nil {
				utils.PushMetrics(scanInfo.GetScanUuid(), msgNo, tstamp, Webhook)
			} else {
				log.Fatal(err)
			}
		}
	} else {
		responses, err := consumers.LoadEnrichedToolResponse()
		if err == nil {
			log.Fatal(err)
		}
		if !LongFormat {
			messages, err := utils.ProcessEnrichedMessages(responses)
			if err == nil {
				log.Fatal(err)
			}
			for _, msg := range messages {
				utils.PushMessage(msg, Webhook)
			}
		} else {
			scanInfo := utils.GetEnrichedScanInfo(responses[0])
			msgNo := utils.CountEnrichedMessages(responses)
			if tstamp, err := ptypes.Timestamp(scanInfo.GetScanStartTime()); err != nil {
				utils.PushMetrics(scanInfo.GetScanUuid(), msgNo, tstamp, Webhook)
			} else {
				log.Fatal(err)
			}
		}
	}
}
