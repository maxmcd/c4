/**
 * @prettier
 * @flow
 */
import config from "./config";
import Control from "./Control";

const tokenErrors = {
    // TODO: include all the ways a token could be invalid
    "token contains an invalid number of segments": true,
    "signature is invalid": true,
};

type ApiResponse = Promise<[{ resp: ?{}, json: ?{} }, ?Error]>;

export default class Api {
    host: string;
    control: Control;
    state: { loginInputValue: string };
    setState: ({}) => void;
    error: ?Error;
    token: ?string;

    constructor(control: Control) {
        this.host = config.ApiHost;
        this.control = control;
    }
    async fetch(path: string, method: string = "GET", body: ?{}): ApiResponse {
        let options: {
            method: string,
            headers: {},
            body?: string,
        } = {
            method: method,
            headers: {
                Accept: "application/json",
                Authorization: this.token || "",
                "Content-Type": "application/json",
            },
        };
        if (body) {
            options.body = JSON.stringify(body);
        }
        let url = `http${config.tls ? "s" : ""}://${this.host}/v0/${path}`;
        // console.log(url);
        let resp = null;
        let json = null;
        try {
            resp = await fetch(url, options);
            json = await resp.json();
            if (resp.status !== 200) {
                if (tokenErrors[json.error]) {
                    this.control.logout();
                }
                console.log(
                    `Request to ${method} ${path} returned ${resp.status}:` +
                        `\n${JSON.stringify(json)}`,
                );
                throw new Error(json.error);
            }
            return [{ resp, json }, null];
        } catch (e) {
            return Promise.resolve([{ resp, json }, e]);
        }
    }
    websocket(path: string, onmessage: ({ data: {} }) => void): WebSocket {
        const url = `ws${config.tls ? "s" : ""}://${this.host}/v0/${path}`;
        console.log("WS connection to ", url);
        let ws = new WebSocket(url, [`Bearer`, `${this.token || ""}`]);
        ws.onopen = () => {
            console.log("Ws open for " + path);
        };
        ws.onmessage = e => {
            if (typeof e.data !== "string") return;
            onmessage(JSON.parse(e.data));
        };
        ws.onerror = e => {
            console.log("Ws error for " + path, e);
        };
        return ws;
    }
    authenticate(body: { phone: string }): ApiResponse {
        return this.fetch("authenticate", "POST", body);
    }
    checkCode(body: {
        phone: string,
        request_id: string,
        code: string,
    }): ApiResponse {
        return this.fetch("authenticate/sms-code-check", "POST", body);
    }
    updateUser(username: string): ApiResponse {
        return this.fetch(`user`, "POST", { username: username });
    }
}
