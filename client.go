package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Client struct {
	AppKey       string
	MasterSecret string
	Domain       string
}

func NewClient(appKey string, masterSecret string) *Client {
	return &Client{
		AppKey:       appKey,
		MasterSecret: masterSecret,
		Domain:       DefaultDomain,
	}
}

func (c *Client) SetDomain(domain string) {
	c.Domain = domain
}

func (c *Client) Broadcast(msg *Message) (*PushResponse, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(body, BroadcastPath, MethodPost)
}

func (c *Client) Unicast(msg *Message, pushId string) (*PushResponse, error) {
	body, err := json.Marshal(&UnicastMessage{Message: msg, PushID: pushId})
	if err != nil {
		return nil, err
	}
	return c.sendRequest(body, UnicastPath, MethodPost)
}

func (c *Client) Muticast(msg *Message, pushIds []string) (*PushResponse, error) {
	body, err := json.Marshal(&MuticastMessage{Message: msg, PushIDs: pushIds})
	if err != nil {
		return nil, err
	}
	return c.sendRequest(body, MuticastPath, MethodPost)
}

func (c *Client) Cuidcast(msg *Message, cuids []string) (*PushResponse, error) {
	body, err := json.Marshal(&CuidsMessage{Message: msg, Cuids: cuids})
	if err != nil {
		return nil, err
	}
	return c.sendRequest(body, CuidsPath, MethodPost)
}

func (c *Client) DelMsg(msgId string) (*PushResponse, error) {
	body, err := json.Marshal(map[string]interface{}{"msg_id": msgId})
	if err != nil {
		return nil, err
	}
	return c.sendRequest(body, DelMsgPath, MethodPost)
}

func (c *Client) sendRequest(body []byte, path string, method string) (*PushResponse, error) {
	var (
		uri      = fmt.Sprintf("%s%s", c.Domain, path)
		ts       = fmt.Sprintf("%d", time.Now().Unix())
		s        = fmt.Sprintf("%s%s%s%s%s%s", method, uri, body, c.AppKey, ts, c.MasterSecret)
		response = &PushResponse{}
	)

	s = url.QueryEscape(s)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(s)))

	args := map[string]string{
		"appkey":    c.AppKey,
		"timestamp": ts,
		"sign":      sign,
	}
	header := map[string]string{
		"Content-Type": "application/json",
	}

	_, b, err := HttpPost(uri, header, args, body)
	if err != nil {
		fmt.Printf("[Error]    %s\n", err)
		return nil, err
	}
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}
	return response, nil
}
