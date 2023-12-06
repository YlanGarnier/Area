import { ScrollView, Text, View, Platform, Alert } from "react-native";
import React, { useEffect, useState } from "react";
import styles from "./style";
import ServiceButton from "../../components/ServiceButton/ServiceButton";
import MyButton from "../../components/Button/Button";
import { NavigationProp, useNavigation } from "@react-navigation/native";
import { fetchTokens } from "../../Functions/tokenUtils";
import * as SecureStore from "expo-secure-store";
import { RootStackParamList } from "../../App";
import LogoGithub from "../../assets/logo-github";
import LogoDiscord from "../../assets/logo-discord";
import LogoFacebook from "../../assets/logo-facebook";
import LogoTwitch from "../../assets/logo-twitch";
import LogoMiro from "../../assets/logo-miro";
import LogoSpotify from "../../assets/logo-spotify";
import LogoNotion from "../../assets/logo-notion";
import LogoDropbox from "../../assets/logo-dropbox";
import LogoLinkedin from "../../assets/logo-linkedin";
import LogoTwitter from "../../assets/logo-twitter";
import LogoGmail from "../../assets/logo-gmail";
import axios from "axios/index";
import * as data from "front-mobile/global.json";

export default function AskService() {
  const navigation = useNavigation<NavigationProp<RootStackParamList>>();
  const [jwtToken, setJwtToken] = useState<string | null>(null);
  const [tokens, setTokens] = useState<{
    googleToken?: string | null;
    google_gmailToken?: string | null;
    discordToken?: string | null;
    facebookToken?: string | null;
    spotifyToken?: string | null;
    miroToken?: string | null;
    githubToken?: string | null;
    twitchToken?: string | null;
    notionToken?: string | null;
    dropboxToken?: string | null;
    linkedinToken?: string | null;
    twitterToken?: string | null;
  }>({});

  function GoUser() {
    navigation.reset({
      index: 0,
      routes: [{ name: "Profile" }],
    });
    navigation.navigate("Lobby", { initialTab: "Profile" });
  }
  useEffect(() => {
    async function initializeTokens() {
      const tokens = await fetchTokens();
      setJwtToken(await SecureStore.getItemAsync("jwtToken"));
      setTokens(tokens);
    }
    initializeTokens();
  }, []);
  const services = [
    {
      onPressLink: `https://accounts.google.com/o/oauth2/v2/auth?client_id=${data.GMAIL_CLIENT_ID}&access_type=offline&response_type=code&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=https://www.googleapis.com/auth/userinfo.email%20https://mail.google.com/&state=google_gmail%20${jwtToken}`,
      logo: <LogoGmail width={40} height={40} />,
      text: "Link Gmail",
      tokenKey: "google_gmailToken",
      tokenValue: tokens.google_gmailToken,
      isLinked: false,
    },
    {
      onPressLink: `https://discord.com/oauth2/authorize?client_id=${data.DISCORD_CLIENT_ID}&redirect_uri=${data.URI_CALLBACK_OAUTH}&response_type=code&scope=identify%20email%20guilds&state=discord%20${jwtToken}`,
      logo: <LogoDiscord width={40} height={40} />,
      text: "Link Discord",
      tokenKey: "discordToken",
      tokenValue: tokens.discordToken,
      isLinked: false,
    },
    {
      onPressLink: `https://www.facebook.com/v18.0/dialog/oauth?client_id=${data.FACEBOOK_CLIENT_ID}&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=email&state=facebook%20${jwtToken}`,
      logo: <LogoFacebook width={40} height={40} />,
      text: "Link Facebook",
      tokenKey: "facebookToken",
      tokenValue: tokens.facebookToken,
      isLinked: false,
    },
    {
      onPressLink: `https://github.com/login/oauth/authorize?client_id=${data.GITHUB_CLIENT_ID_MOBILE_IOS}&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=repo%20user&state=github%20${jwtToken}`,
      logo: <LogoGithub width={40} height={40} />,
      text: "Link Github",
      tokenKey: "githubToken",
      tokenValue: tokens.githubToken,
      isLinked: false,
    },
    {
      onPressLink: `https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=${data.TWITCH_CLIENT_ID}&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=user%3Aedit%20user%3Aread%3Aemail%20channel%3Amanage%3Apolls%20moderator%3Amanage%3Aannouncements&state=twitch%20${jwtToken}`,
      logo: <LogoTwitch width={40} height={40} />,
      text: "Link Twitch",
      tokenKey: "twitchToken",
      tokenValue: tokens.twitchToken,
      isLinked: false,
    },
    {
      onPressLink: `https://accounts.spotify.com/authorize?client_id=${data.SPOTIFY_CLIENT_ID}&response_type=code&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=user-read-email%2Cuser-read-currently-playing&state=spotify%20${jwtToken}`,
      logo: <LogoSpotify width={40} height={40} />,
      text: "Link Spotify",
      tokenKey: "spotifyToken",
      tokenValue: tokens.spotifyToken,
      isLinked: false,
    },
    {
      onPressLink: `https://miro.com/oauth/authorize?client_id=${data.MIRO_CLIENT_ID}&redirect_uri=${data.URI_CALLBACK_OAUTH}&response_type=code&scope=email&state=miro%20${jwtToken}`,
      logo: <LogoMiro width={40} height={40} />,
      text: "Link Miro",
      tokenKey: "miroToken",
      tokenValue: tokens.miroToken,
      isLinked: false,
    },
    {
      onPressLink: `https://api.notion.com/v1/oauth/authorize?client_id=${data.NOTION_CLIENT_ID}&response_type=code&owner=user&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=email&state=notion%20${jwtToken}`,
      logo: <LogoNotion width={40} height={40} />,
      text: "Link Notion",
      tokenKey: "notionToken",
      tokenValue: tokens.notionToken,
      isLinked: false,
    },
    {
      onPressLink: `https://www.dropbox.com/oauth2/authorize?client_id=${data.DROPBOX_CLIENT_ID}&redirect_uri=${data.URI_CALLBACK_OAUTH}&response_type=code&state=dropbox%20${jwtToken}`,
      logo: <LogoDropbox width={40} height={40} />,
      text: "Link Dropbox",
      tokenKey: "dropboxToken",
      tokenValue: tokens.dropboxToken,
      isLinked: false,
    },
    {
      onPressLink: `https://www.linkedin.com/oauth/v2/authorization?response_type=code&client_id=${data.LINKEDIN_CLIENT_ID}&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=profile%20email%20openid%20w_member_social&state=linkedin%20${jwtToken}`,
      logo: <LogoLinkedin width={40} height={40} />,
      text: "Link Linkedin",
      tokenKey: "linkedinToken",
      tokenValue: tokens.linkedinToken,
      isLinked: false,
    },
    {
      onPressLink: `https://twitter.com/i/oauth2/authorize?client_id=${data.TWITTER_CLIENT_ID}&scope=tweet.read%20users.read%20follows.read%20follows.write%20tweet.write&code_challenge=challenge&code_challenge_method=plain&redirect_uri=${data.URI_CALLBACK_OAUTH}&response_type=code&state=twitter%20${jwtToken}`,
      logo: <LogoTwitter width={40} height={40} />,
      text: "Link Twitter",
      tokenKey: "twitterToken",
      tokenValue: tokens.twitterToken,
      isLinked: false,
    },
  ];

  const [linkedServices, setLinkedServices] = useState([]);

  useEffect(() => {
    const fetchToken = async () => {
      try {
        const token = await SecureStore.getItemAsync("jwtToken");
        const headers = { Authorization: `Bearer ${token}` };
        const response = await axios.get(`${data.API_URL}/users/me/services`, {
          headers: headers,
        });
        const updatedServices = services.map((service) => {
          for (let j = 0; j < response.data.length; j++) {
            if (response.data[j].name === service.tokenKey.slice(0, -5)) {
              return { ...service, isLinked: true };
            }
          }
          return service;
        });

        setLinkedServices(updatedServices);
      } catch (error) {
        Alert.alert("ERROR == ", error);
      }
    };

    fetchToken();
  }, [linkedServices]);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Connect to different services :</Text>
      <ScrollView style={styles.scroll}>
        <View style={styles.linksContainer}>
          {linkedServices.map((service, index) => (
            <ServiceButton
              key={index}
              onPressLink={service.onPressLink}
              logo={service.logo}
              tokenKey={service.tokenKey}
              tokenValue={service.tokenValue}
              isLinked={service.isLinked}
              text={service.text}
              style={styles.links}
              textStyle={styles.linkText}
            />
          ))}
        </View>
      </ScrollView>
      <MyButton
        name={"Go to user page"}
        ButtonStyle={styles.finish}
        TextStyle={styles.finishText}
        onPress={GoUser}
      />
    </View>
  );
}
