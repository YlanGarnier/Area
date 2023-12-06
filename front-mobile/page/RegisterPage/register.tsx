import { useState } from "react";
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
import * as data from "front-mobile/global.json";
import { StackScreenProps } from "@react-navigation/stack";
import * as SecureStore from "expo-secure-store";

type RootStackParamList = {
  Welcome: undefined;
  Register: undefined;
  RegisterInfo: undefined;
  AskService: undefined;
  Lobby: undefined;
  Area: undefined;
  Home: undefined;
};

export default function Register({
  navigation,
}: StackScreenProps<RootStackParamList, "Register">) {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confPass, setConfPass] = useState("");
  const [message, setMessage] = useState("");

  function isValidEmail(email: string) {
    const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
    return emailRegex.test(email);
  }
  function isValidPassword(email: string) {
    const passwordRegex =
      /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$/;
    return passwordRegex.test(password);
  }
  const onSubmitFormHandler = async () => {
    if (password != confPass) {
      setMessage("Passwords does not match");
      return;
    }
    if (!isValidEmail(email)) {
      setMessage("Invalid email");
      return;
    }
    if (!isValidPassword(password)) {
      setMessage("Password too weak");
      return;
    }
    Alert.alert(`${data.API_URL}/users/signup`);
    axios
      .post(`${data.API_URL}/users/signup`, { username, email, password })
      .then(async (response) => {
        if (!response.data?.status) {
          await SecureStore.setItemAsync("jwtToken", response.data.token);
          navigation.reset({
            index: 0,
            routes: [{ name: "RegisterInfo" }],
          });
        } else {
          setMessage("error invalid credentials");
        }
      })
      .catch(() => {
        setMessage("error on request");
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
          <View style={styles.registerContainer}>
            <Text style={styles.registerText}>Already have an account? </Text>
            <TouchableOpacity onPress={() => navigation.navigate("Welcome")}>
              <Text style={styles.signupText}>Sign In</Text>
            </TouchableOpacity>
          </View>
          <Text style={styles.catchphrase}>
            Please fill your information below
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
          <View style={styles.inputContainer}>
            <Text style={styles.label}>Confirm password</Text>
            <TextInput
              onChangeText={setConfPass}
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
            <Text style={styles.signinText}>SIGN UP</Text>
          </TouchableOpacity>
        </View>
        <StatusBar style="auto" />
      </View>
    </KeyboardAvoidingView>
  );
}
