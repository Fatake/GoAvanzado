package main

import "fmt"

// Notificaciones SMS E EMAIL

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotificacion struct {
}

//Para que SMSNotificacion implemente una interfaces
// tenemos que hacer un metodo con el nombre exacto de alguna interface

func (SMSNotificacion) SendNotification() {
	fmt.Print("[+] enviado notificacion via Sms")
}

func (SMSNotificacion) GetSender() ISender {
	return SMSNotificationSender{}
}

// Imprementacion del ISender
type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "WhatssApp"
}
