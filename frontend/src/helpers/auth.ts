import { API } from "../api/api.ts";
import {
  AuthenticationResponseJSON,
  RegistrationResponseJSON,
} from "@simplewebauthn/typescript-types";
import {
  startAuthentication,
  startRegistration,
} from "@simplewebauthn/browser";
import { Router } from "vue-router";
import { QueryClient } from "@tanstack/vue-query";

export const tryLogin = async (router: Router, queryClient: QueryClient) => {
  const numCreds = await API.auth.checkRegistration();
  let token: string;
  if (numCreds === 0) {
    token = await register();
  } else {
    token = await login();
  }
  document.cookie = `auth=${token};max-age=3600`;
  await queryClient.invalidateQueries({ queryKey: ["login"] });
  await router.push("/");
};

const register = async () => {
  const beginRegistrationResp = await API.auth.beginRegistration();
  let attResp: RegistrationResponseJSON;
  try {
    attResp = await startRegistration(beginRegistrationResp.publicKey);
  } catch (error) {
    console.log(error);
    throw error;
  }

  return await API.auth.finishRegistration(attResp);
};

const login = async () => {
  const beginLoginResp = await API.auth.beginLogin();
  let attResp: AuthenticationResponseJSON;
  try {
    attResp = await startAuthentication(beginLoginResp.publicKey);
  } catch (error) {
    console.log(error);
    throw error;
  }

  return await API.auth.finishLogin(attResp);
};
