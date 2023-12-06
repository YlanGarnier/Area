import { writable } from 'svelte/store';

export const authToken = writable<string | null>(localStorage.getItem('authToken'));

export const pageTitle = writable<string | null>(null);

interface ValidString {
  String: string;
  Valid: boolean;
}

interface UserInfo {
  Email: string;
  FirstName: ValidString;
  Kind: string;
  LastName: ValidString;
  ID: number;
  Username: ValidString;
}

export const userInfos = writable<UserInfo | null>(null);

interface ActionTypes {
  service: string,
  route: string,
}

export const Action = writable<ActionTypes | null>(null);

export const Reaction = writable<ActionTypes | null>(null);

interface ServicesTypes {
  success: string | null,
  failed: string | null
}

export const servicesStorage = writable<ServicesTypes | null>(null);

interface AreaInterface {
  action_service: string,
  id: number,
  name: string,
  reaction_service: string,
  route_action_service: string,
  route_reaction_service: string
}

export const AreaStorage = writable<AreaInterface | null>(null);

export const AuthStatus = writable<boolean>(false);
