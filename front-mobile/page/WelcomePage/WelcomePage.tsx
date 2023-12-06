import { useEffect, useState } from "react";
import {
  KeyboardAvoidingView,
  Platform,
  Image,
  Text,
  TextInput,
  TouchableOpacity,
  View,
  Alert,
} from "react-native";
import { StatusBar } from "expo-status-bar";
import styles from "./style";
import axios from "axios";
import * as SecureStore from "expo-secure-store";
import * as data from "front-mobile/global.json";
import { StackScreenProps } from "@react-navigation/stack";
import LogoGithub from "../../assets/logo-github";
import LogoGoogle from "../../assets/logo-google";
import LogoDiscord from "../../assets/logo-discord";
import * as Linking from "expo-linking";
import Redirection from "../../Functions/Redirection";
import {
  fetchAndStoreAllServiceTokens,
  fetchTokens,
  sendToken,
} from "../../Functions/tokenUtils";

type RootStackParamList = {
  Welcome: undefined;
  Register: undefined;
  RegisterInfo: undefined;
  Lobby: undefined;
  Area: undefined;
  Home: undefined;
};

export default function WelcomePage({
  navigation,
}: StackScreenProps<RootStackParamList, "Welcome">) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleURL = async (event: { url: string }) => {
    const { path, queryParams } = Linking.parse(event.url);

    if (!queryParams.state) {
      return;
    }
    const { state } = queryParams;
    if (state === "google" || state === "discord") {
      await SecureStore.setItemAsync("jwtToken", queryParams.token);
      if (queryParams.status === "302") {
        navigation.navigate("Lobby");
      } else {
        navigation.navigate("RegisterInfo");
      }
    }
  };

  useEffect(() => {
    const handleDeepLink = (event: { url: string }) => {
      handleURL(event);
    };
    const subscription = Linking.addEventListener("url", handleDeepLink);
    return () => {
      subscription.remove();
    };
  }, []);

  const onSubmitFormHandler = async () => {
    axios
      .post(`${data.API_URL}/users/login`, { email, password })
      .then(async (response) => {
        if (!response.data?.status) {
          if (response.data.token) {
            await SecureStore.setItemAsync("jwtToken", response.data.token);
            await fetchAndStoreAllServiceTokens();
            await fetchTokens();
          } else {
            Alert.alert("JWTToken is not valid:", response.data.token);
          }
          navigation.reset({
            index: 0,
            routes: [{ name: "Lobby" }],
          });
        } else {
          setMessage("error invalid credentials");
        }
      })
      .catch(() => {
        setMessage("error invalid credentials");
      });
  };

  return (
    <KeyboardAvoidingView
      style={styles.container}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      enabled
    >
      <View style={styles.container}>
        <Image
          style={styles.background}
          source={require("front-mobile/assets/HomeBg.jpeg")}
        />
        <View style={styles.middle}>
          <Text style={styles.catchphrase}>
            Please enter your credentials below
          </Text>
          <View style={styles.inputContainer}>
            <Text style={styles.label}>Email</Text>
            <TextInput onChangeText={setEmail} style={styles.input} />
          </View>
          <View style={styles.inputContainer}>
            <Text style={styles.label}>Password</Text>
            <TextInput
              onChangeText={setPassword}
              style={styles.input}
              secureTextEntry={true}
            />
          </View>
          <View style={styles.messageContainer}>
            <Text style={styles.errmessage}>{message}</Text>
          </View>
          <TouchableOpacity
            style={styles.signButton}
            onPress={onSubmitFormHandler}
          >
            <Text style={styles.signinText}>LOGIN</Text>
          </TouchableOpacity>
          <View style={styles.registerContainer}>
            <Text style={styles.registerText}>Don't have an account? </Text>
            <TouchableOpacity onPress={() => navigation.navigate("Register")}>
              <Text style={styles.signupText}>Sign Up</Text>
            </TouchableOpacity>
          </View>
        </View>
        <View style={styles.dividerContainer}>
          <View style={styles.dividerLine} />
          <Text style={styles.dividerText}>or continue with</Text>
          <View style={styles.dividerLine} />
        </View>
        <View style={styles.socialContainer}>
          <TouchableOpacity
            style={styles.social}
            onPress={() =>
              Redirection(
                `https://accounts.google.com/o/oauth2/v2/auth?client_id=${data.GOOGLE_CLIENT_ID}&access_type=offline&response_type=code&redirect_uri=${data.URI_CALLBACK_OAUTH}&scope=https://www.googleapis.com/auth/userinfo.email&state=google`,
              )
            }
          >
            <LogoGoogle width={60} height={60} />
          </TouchableOpacity>
          <TouchableOpacity
            style={styles.social}
            onPress={() =>
              Redirection(
                `https://discord.com/oauth2/authorize?client_id=${data.DISCORD_CLIENT_ID}&redirect_uri=${data.URI_CALLBACK_OAUTH}&response_type=code&scope=identify%20email&state=discord`,
              )
            }
          >
            <LogoDiscord width={60} height={60} />
          </TouchableOpacity>
        </View>
        <StatusBar style="auto" />
      </View>
    </KeyboardAvoidingView>
  );
}
