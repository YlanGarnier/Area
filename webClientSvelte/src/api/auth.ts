import axios from 'axios';
import { navigate } from "svelte-routing";
import { get } from 'svelte/store';
import { AuthStatus, authToken, userInfos } from './stores';

const API_URL: string = import.meta.env.VITE_API_URL;
const OAUTH_REDIRECT_URI: string = import.meta.env.VITE_OAUTH_REDIRECT_URI;
const SERVICES_REDIRECT_URI: string = import.meta.env.VITE_SERVICES_REDIRECT_URI;

function isValidEmail(email: string): boolean {
  const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
  return emailRegex.test(email);
}

export async function userLogin(email: string, password: string, message: string) {
  if (!isValidEmail(email)) {
    message = "Invalid email";
    return;
  }

  try {
    const response = await axios.post(`${API_URL}/users/login`, { email, password });
    if (response.data && response.status === 200) {
      message = "Signed in successfully";
      authToken.set(response.data.token)
      navigate("/home");
    } else {
      message = "Error: Invalid credentials";
    }
  } catch (error) {
    message = "Error on request";
  }
}

export async function userSignup(email: string, password: string, confPassword: string): Promise<string> {
  if (password !== confPassword) {
    return "Passwords do not match";
  }
  if (!isValidEmail(email)) {
    return "Invalid email";
  }

  try {
    const response = await axios.post(`${API_URL}/users/signup`, { email, password });
    return "Signed up successfully";
  } catch (error) {
    console.log(error);
    return "Error on request";
  }
}

export async function oAuthSignup(code: string, platform: string, provider: string) {
  try {
    const response = await axios.post(`${API_URL}/users/signup/oauth`, { code, platform, provider, redirect_uri: OAUTH_REDIRECT_URI }, { maxRedirects: 0 });
    AuthStatus.set(false);
    return response.data.token;
  } catch (error: any) {
    if (error.response.status === 302) {
      AuthStatus.set(true);
      return error.response.data.token;
    }
    else if (error.response)
      throw new Error(error.response.data.message || "Failed to retrieve user infos.");
    else if (error.request)
      throw new Error("No response received from the server.");
    else
      throw error;
  }
}

export async function serviceConnect(code: string, platform: string, provider: string) {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    const headers = { "Authorization": `Bearer ${token}` };
    console.log(code, platform, provider, SERVICES_REDIRECT_URI);
    const response = await axios.post(`${API_URL}/users/me/services`, { code, platform, provider, redirect_uri: SERVICES_REDIRECT_URI }, { headers });

    if (response.status === 201) {
      return;
    } else {
      throw new Error("Connection to service failed.");
    }
  } catch (error: any) {
    if (error.response) {
      // The request was made and the server responded with a status code outside of the 2xx range
      throw new Error(error.response.data.message || "Failed to retrieve user infos.");
    } else if (error.request) {
      // The request was made but no response was received
      throw new Error("No response received from the server.");
    } else {
      // Something happened in setting up the request and triggered an error
      throw error;
    }
  }
}

interface AvailableServices {
  name: string
}

export async function getServices(): Promise<AvailableServices[]> {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    const headers = { "Authorization": `Bearer ${token}` };
    const response = await axios.get<AvailableServices[]>(`${API_URL}/users/me/services`, { headers });
    return response.data;
  } catch (error: any) {
    if (error.response) {
      // The request was made and the server responded with a status code outside of the 2xx range
      throw new Error(error.response.data.message || "Failed to retrieve user infos.");
    } else if (error.request) {
      // The request was made but no response was received
      throw new Error("No response received from the server.");
    } else {
      // Something happened in setting up the request and triggered an error
      throw error;
    }
  }
}

export async function newArea(action: any, reaction: any, name: string) {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token");

    const headers = { "Authorization": `Bearer ${token}` };
    const payload = {
      action: {
        service: action.name,
        route: action.route,
        params: action.params
      },
      name: name,
      reaction: {
        service: reaction.name,
        route: reaction.route,
        target: reaction.params[0]
      },
    };

    await axios.post(`${API_URL}/area/new`, payload, { headers })
    return;
  } catch (error: any) {
    if (error.response) {
      throw new Error(error.response.data.message || "Failed to create the area.");
    } else if (error.request) {
      throw new Error("No response received from the server.");
    } else {
      throw error;
    }
  }
}

interface AreaInterface {
  action_service: string,
  id: number,
  name: string,
  reaction_service: string,
  route_action_service: string,
  route_reaction_service: string
}

export async function getAreas(): Promise<AreaInterface[]> {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token");

    const headers = { "Authorization": `Bearer ${token}` };
    const response = await axios.get<AreaInterface[]>(`${API_URL}/users/me/areas`, { headers });
    return response.data;
  } catch (error: any) {
    if (error.response) {
      throw new Error(error.response.data.message || "Failed to create the area.");
    } else if (error.request) {
      throw new Error("No response received from the server.");
    } else {
      throw error;
    }
  }
}

export async function deleteArea(id: number) {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token");

    const headers = { "Authorization": `Bearer ${token}` };
    await axios.delete(`${API_URL}/area/${id}`, { headers })
    return;
  } catch (error: any) {
    if (error.response) {
      throw new Error(error.response.data.message || "Failed to create the area.");
    } else if (error.request) {
      throw new Error("No response received from the server.");
    } else {
      throw error;
    }
  }
}

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

export async function getUserInfo() {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    const headers = { "Authorization": `Bearer ${token}` };
    const response = await axios.get<UserInfo>(`${API_URL}/users/me`, { headers });
    userInfos.set(response.data);
  } catch (error: any) {
    if (error.response) {
      // The request was made and the server responded with a status code outside of the 2xx range
      throw new Error(error.response.data.message || "Failed to retrieve user infos.");
    } else if (error.request) {
      // The request was made but no response was received
      throw new Error("No response received from the server.");
    } else {
      // Something happened in setting up the request and triggered an error
      throw error;
    }
  }
}

export async function updateEmail(email: string, password: string) {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    const headers = { "Authorization": `Bearer ${token}` };
    await axios.put(`${API_URL}/users/me/email`, { email: email, password: password }, { headers });
  } catch (error: any) {
    if (error.response) {
      // The request was made and the server responded with a status code outside of the 2xx range
      throw new Error(error.response.data.message || "Failed to retrieve user infos.");
    } else if (error.request) {
      // The request was made but no response was received
      throw new Error("No response received from the server.");
    } else {
      // Something happened in setting up the request and triggered an error
      throw error;
    }
  }
}

export async function updatePassword(password: string, newPassword: string) {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    const headers = { "Authorization": `Bearer ${token}` };
    await axios.put(`${API_URL}/users/me/password`, { new_password: newPassword, previous_password: password }, { headers });
  } catch (error: any) {
    if (error.response) {
      // The request was made and the server responded with a status code outside of the 2xx range
      throw new Error(error.response.data.message || "Failed to retrieve user infos.");
    } else if (error.request) {
      // The request was made but no response was received
      throw new Error("No response received from the server.");
    } else {
      // Something happened in setting up the request and triggered an error
      throw error;
    }
  }
}

export async function updateInfos(firstName?: string, lastName?: string, userName?: string) {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    const headers = { "Authorization": `Bearer ${token}` };
    const body: Record<string, string> = {};
    if (firstName !== undefined) body.first_name = firstName;
    if (lastName !== undefined) body.last_name = lastName;
    if (userName !== undefined) body.username = userName;
    if (Object.keys(body).length === 0) {
      throw new Error("No update information provided.");
    }

    await axios.put(`${API_URL}/users/me`, body, { headers });
  } catch (error: any) {
    if (error.response) {
      // The request was made and the server responded with a status code outside of the 2xx range
      throw new Error(error.response.data.message || "Failed to retrieve user infos.");
    } else if (error.request) {
      // The request was made but no response was received
      throw new Error("No response received from the server.");
    } else {
      // Something happened in setting up the request and triggered an error
      throw error;
    }
  }
}


export async function logout() {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    authToken.set(null);
    navigate("/");
  } catch (error) {
    throw error;
  }
}

export async function deleteAccount() {
  try {
    const token = get(authToken);
    if (!token)
      throw new Error("Failed to retrieve token.");

    const headers = { "Authorization": `Bearer ${token}` };
    await axios.delete(`${API_URL}/users/me`, { headers });
  } catch (error: any) {
    if (error.response) {
      throw new Error(error.response.data.message || "Failed to create the area.");
    } else if (error.request) {
      throw new Error("No response received from the server.");
    } else {
      throw error;
    }
  }
}
