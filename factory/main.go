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
	fmt.Println("[+] enviado notificacion via Sms")
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

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("[+] Enviando por Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Hotmail"
}

func getNotificationFactory(notificacion string) (INotificationFactory, error) {
	if notificacion == "SMS" {
		return &SMSNotificacion{}, nil
	}
	if notificacion == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("[!] No se encuentra el tipo de notificacion")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	fmt.Println("[i] Iniciando programa de NotificationFactory")
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")
	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
