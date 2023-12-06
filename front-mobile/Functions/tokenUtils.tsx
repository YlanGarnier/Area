import * as SecureStore from "expo-secure-store";
import axios from "axios";
import * as data from "front-mobile/global.json";
import Alert from "react-native";

export async function initToken() {
  await SecureStore.setItemAsync("jwtToken", "");
  await SecureStore.setItemAsync("githubToken", "");
  await SecureStore.setItemAsync("google_gmailToken", "");
  await SecureStore.setItemAsync("discordToken", "");
  await SecureStore.setItemAsync("googleToken", "");
  await SecureStore.setItemAsync("spotifyToken", "");
  await SecureStore.setItemAsync("facebookToken", "");
  await SecureStore.setItemAsync("twitchToken", "");
  await SecureStore.setItemAsync("miroToken", "");
  await SecureStore.setItemAsync("notionToken", "");
  await SecureStore.setItemAsync("dropboxToken", "");
  await SecureStore.setItemAsync("linkedinToken", "");
  await SecureStore.setItemAsync("twitterToken", "");
}

export async function fetchTokens() {
  const google_gmailToken = await SecureStore.getItemAsync("google_gmailToken");
  const googleToken = await SecureStore.getItemAsync("googleToken");
  const discordToken = await SecureStore.getItemAsync("discordToken");
  const facebookToken = await SecureStore.getItemAsync("facebookToken");
  const spotifyToken = await SecureStore.getItemAsync("spotifyToken");
  const miroToken = await SecureStore.getItemAsync("miroToken");
  const githubToken = await SecureStore.getItemAsync("githubToken");
  const twitchToken = await SecureStore.getItemAsync("twitchToken");
  const notionToken = await SecureStore.getItemAsync("notionToken");
  const dropboxToken = await SecureStore.getItemAsync("dropboxToken");
  const linkedinToken = await SecureStore.getItemAsync("linkedinToken");
  const twitterToken = await SecureStore.getItemAsync("twitterToken");

  return {
    googleToken,
    google_gmailToken,
    discordToken,
    facebookToken,
    spotifyToken,
    miroToken,
    githubToken,
    twitchToken,
    notionToken,
    dropboxToken,
    linkedinToken,
    twitterToken,
  };
}

export const fetchAndStoreAllServiceTokens = async () => {
  const headers = {
    Authorization: `Bearer ${await SecureStore.getItemAsync("jwtToken")}`,
  };
  try {
    const response = await axios.get(`${data.API_URL}/users/me/services`, {
      headers,
    });
    if (response.data && Array.isArray(response.data)) {
      for (const service of response.data) {
        const tokenKey = `${service.name}Token`;
        await SecureStore.setItemAsync(tokenKey, "set");
      }
    }
  } catch (error) {
    Alert.alert("Error while fetching service tokens:", error);
  }
};

export async function sendToken(
  code: string,
  platform: string,
  provider: string,
  redirect_uri: string,
) {
  try {
    const headers = {
      Authorization: `Bearer ${await SecureStore.getItemAsync("jwtToken")}`,
    };
    const response = await axios.post(
      `${data.API_URL}/users/me/services`,
      { code, platform, provider, redirect_uri },
      { headers },
    );
  } catch (e) {
    Alert.alert(e);
  }
}

export async function deleteToken(tokenName: string) {
  await SecureStore.deleteItemAsync(tokenName);
}

export async function deleteAllTokens() {
  const tokenKeys = [
    "jwtToken",
    "googleToken",
    "google_gmailToken",
    "discordToken",
    "facebookToken",
    "spotifyToken",
    "miroToken",
    "githubToken",
    "twitchToken",
    "notionToken",
    "dropboxToken",
    "linkedinToken",
    "twitterToken",
  ];
  for (const key of tokenKeys) {
    await SecureStore.deleteItemAsync(key);
  }
}
