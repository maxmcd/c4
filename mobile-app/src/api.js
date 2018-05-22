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

type fetchOptions = {
    method: string,
    headers: {},
    body?: string,
};

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
    fetchWTimeout(url: string, options: fetchOptions) {
        let promise = fetch(url, options);
        return new Promise(function(resolve, reject) {
            promise.then(resolve, reject);
            // TODO: this doesn't cancel the request
            // likely switch to XMLHttpRequest
            // https://facebook.github.io/react-native/docs/network.html
            // https://github.com/matthew-andrews/isomorphic-fetch/issues/48#issuecomment-380465212
            setTimeout(reject.bind(null, new Error("Request timed out")), 5000);
        });
    }
    async fetch(path: string, method: string = "GET", body: ?{}): ApiResponse {
        let options: fetchOptions = {
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
        console.log(url);
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
    updateUser(username: string): ApiResponse {
        return this.fetch(`user`, "POST", { username: username });
    }
    joinPool(): ApiResponse {
        return this.fetch(`game/join-pool`, "POST", {});
    }
    sendMove(gameId: string, command: string, column: number): ApiResponse {
        return this.fetch(`game/${gameId}/send-move`, "POST", {
            column: column,
            command: command,
        });
    }
    receiveMoves(gameId: string, onmessage: ({ data: {} }) => void): WebSocket {
        return this.websocket(`game/${gameId}/recevie-move`, onmessage);
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
    checkCode(body: { phone: string, code: string }): ApiResponse {
        return this.fetch("authenticate/sms-code-check", "POST", body);
    }
}
