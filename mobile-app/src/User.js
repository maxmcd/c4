/**
 * @prettier
 * @flow
 */
import Api from "./api";
import Control from "./Control";
import Base64 from "./Base64";
import config from "./config";
import { AsyncStorage } from "react-native";

type UserSave = {
  token: string,
  user_id: string,
  username: string,
};

export default class User {
  api: Api;
  userId: string;
  token: ?string;
  username: ?string;
  email: ?string;
  code: ?string;
  phoneNumber: ?string;

  constructor(control: Control) {
    this.api = new Api(control);
    this.token = null;
    this.code = null;
    this.phoneNumber = null;
  }
  async loadUser(): Promise<?UserSave> {
    try {
      const userSaveString = await AsyncStorage.getItem(
        config.UserSaveStoreKey,
      );
      console.log(userSaveString);
      if (userSaveString) {
        let userSave: UserSave = JSON.parse(userSaveString);
        this.setToken(userSave.token);
        this.username = userSave.username;
        return userSave;
      }
    } catch (e) {
      return console.log(e);
    }
  }
  async saveUser(user: UserSave) {
    try {
      await AsyncStorage.setItem(config.UserSaveStoreKey, JSON.stringify(user));
    } catch (e) {
      console.log(`Error setting jwt: ${e}`);
    }
  }
  async authenticate(phoneNumber: string): Promise<?Error> {
    this.phoneNumber = phoneNumber;
    let apiResp = await this.api.authenticate({
      phone: phoneNumber,
    });
    if (apiResp instanceof Promise) {
      apiResp = await apiResp;
    }
    const [, err] = apiResp;
    if (err) {
      return err;
    }
    return null;
  }
  setToken(token: string): void {
    this.userId = JSON.parse(Base64.atob(token.split(".")[1])).user_id;
    this.api.token = token;
    this.token = token;
  }
  async checkCode(code: string): Promise<?Error> {
    if (!this.phoneNumber) {
      return new Error("Need phone number with checkCode");
    }
    let apiResp = await this.api.checkCode({
      phone: this.phoneNumber,
      code: code,
    });
    if (apiResp instanceof Promise) {
      apiResp = await apiResp;
    }
    const [resp, err] = apiResp;
    if (err) {
      return err;
    }
    this.userId = resp.json.user_id;
    this.username = resp.json.username;
    this.setToken(resp.json.token);
    await this.saveUser(resp.json);
    return null;
  }
  async updateUser(username: string): Promise<?Error> {
    let apiResp = await this.api.updateUser(username);
    if (apiResp instanceof Promise) {
      apiResp = await apiResp;
    }
    const [, err] = apiResp;
    if (err) {
      return err;
    }
    this.username = username;
    return null;
  }
  logout() {
    AsyncStorage.removeItem(config.UserSaveStoreKey).catch(error => {
      console.log(`Error unsetting jwt: ${error}`);
    });
  }
}
