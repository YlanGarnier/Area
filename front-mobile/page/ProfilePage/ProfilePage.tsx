import {
  Text,
  View,
  Modal,
  Alert,
  TouchableOpacity,
  TextInput,
} from "react-native";
import { useEffect, useState, useRef } from "react";
import styles from "./style";
import MyButton from "../../components/Button/Button";
import * as SecureStore from "expo-secure-store";
import { useNavigation, NavigationProp } from "@react-navigation/native";
import { StackScreenProps } from "@react-navigation/stack";
import axios from "axios";
import * as data from "front-mobile/global.json";
import { RootStackParamList } from "../../App";
import { HeaderComponent } from "../../components/Header/HeaderComponent";
import { HeaderButtonComponent } from "../../components/Header/HeaderButtonComponent";
import { Ionicons } from "@expo/vector-icons";
import { LinearGradient } from "expo-linear-gradient";
import { deleteAllTokens } from "../../Functions/tokenUtils";

export default function ProfilePage({
  route,
}: StackScreenProps<RootStackParamList, "Profile">) {
  const navigation = useNavigation<NavigationProp<RootStackParamList>>();
  const [firstName, setFirstName] = useState<string>("");
  const [lastName, setLastName] = useState<string>("");
  const [areaCount, setAreaCount] = useState<number>(2);
  const [email, setEmail] = useState<string>("");
  const [username, setUsername] = useState<string>("");
  const [newPassword, setNewPassword] = useState<string>("");
  const [previousPassword, setPreviousPassword] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [verificationPassword, setVerificationPassword] = useState<string>("");
  const [kind, setKind] = useState<string>("");
  const [isBottomModalVisible, setIsBottomModalVisible] = useState(false);
  const [isEditProfileModalVisible, setIsEditProfileModalVisible] =
    useState(false);

  const initialProfileState = useRef({
    firstName: "",
    lastName: "",
    email: "",
    username: "",
    verificationPassword: "",
  });

  const [tempProfile, setTempProfile] = useState({
    firstName: "",
    lastName: "",
    email: "",
    username: "",
    newPassword: "",
    previousPassword: "",
    verificationPassword: "",
  });

  const openEditProfileModal = () => {
    setTempProfile({
      firstName,
      lastName,
      email,
      username,
      newPassword: "",
      previousPassword: "",
      verificationPassword: "",
    });
    setIsEditProfileModalVisible(true);
  };

  const closeEditProfileModalWithoutSaving = () => {
    setTempProfile(initialProfileState.current);
    setIsEditProfileModalVisible(false);
  };

  const hasProfileChanged = () => {
    return (
      initialProfileState.current.firstName !== tempProfile.firstName ||
      initialProfileState.current.lastName !== tempProfile.lastName ||
      initialProfileState.current.email !== tempProfile.email ||
      initialProfileState.current.username !== tempProfile.username
    );
  };

  async function GetInfo() {
    try {
      const headers = {
        Authorization: `Bearer ${await SecureStore.getItemAsync("jwtToken")}`,
      };
      const response = await axios.get(`${data.API_URL}/users/me`, { headers });

      if (!response.data?.status) {
        setUsername(response.data.Username.String || "undefined");
        setEmail(response.data.Email || "undefined");
        setFirstName(response.data.FirstName.String || "undefined");
        setLastName(response.data.LastName.String || "undefined");
        setKind(response.data.Kind || "undefined");

        initialProfileState.current = {
          firstName: response.data.FirstName || "undefined",
          lastName: response.data.LastName || "undefined",
          email: response.data.Email || "undefined",
          username: response.data.Username || "undefined",
          verificationPassword: "",
        };
      } else {
        Alert.alert("Error: can't get info");
      }
    } catch (error) {
      Alert.alert("Error: can't get info");
    }
  }

  async function Disconnection() {
    try {
      await deleteAllTokens();
      setIsBottomModalVisible(false);
      navigation.reset({
        index: 0,
        routes: [{ name: "Welcome" }],
      });
    } catch (error) {
      Alert.alert("Failed to delete token:");
    }
  }

  async function DeleteAccount() {
    try {
      const headers = {
        Authorization: `Bearer ${await SecureStore.getItemAsync("jwtToken")}`,
      };
      const response = await axios.delete(`${data.API_URL}/users/me`, {
        headers,
      });

      if (!response.data?.status) {
        setIsBottomModalVisible(false);
        await deleteAllTokens();
        navigation.reset({
          index: 0,
          routes: [{ name: "Welcome" }],
        });
      } else {
        Alert.alert("Error while deleting your account");
      }
    } catch (error) {
      Alert.alert("Error during account deletion:");
    }
  }

  useEffect(() => {
    GetInfo();
  }, []);

  function navigateToAskService() {
    setIsBottomModalVisible(false);
    navigation.navigate("AskService", { fromRegister: false });
  }

  function isValidEmail(email: string) {
    const emailRegex = /^[a-zA-Z0-9._-]+@[a-zAZ0-9.-]+\.[a-zA-Z]{2,6}$/;
    return emailRegex.test(email);
  }

  function isValidPassword(password: string) {
    const passwordRegex =
      /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$/;
    return passwordRegex.test(password);
  }

  async function saveChanges() {
    const token = await SecureStore.getItemAsync("jwtToken");
    const headers = { Authorization: `Bearer ${token}` };

    if (!hasProfileChanged()) {
      setIsEditProfileModalVisible(false);
      return;
    }

    let errors = [];
    if (tempProfile.username === "") errors.push("Username");
    if (tempProfile.email === "") errors.push("Email");
    if (tempProfile.firstName === "") errors.push("First Name");
    if (tempProfile.lastName === "") errors.push("Last Name");
    if (hasEmailChanged() && tempProfile.verificationPassword === "")
      errors.push("Confirm Password");
    if (
      kind === "password" &&
      (tempProfile.newPassword || tempProfile.previousPassword)
    ) {
      if (
        tempProfile.newPassword &&
        !isValidPassword(tempProfile.newPassword)
      ) {
        errors.push("New Password is invalid.");
      }
      if (!tempProfile.newPassword) errors.push("New Password is required.");
      if (!tempProfile.previousPassword)
        errors.push("Previous Password is required.");
    }
    if (errors.length > 0) {
      Alert.alert(
        "Error",
        `Please respect the following indications:\n${errors.join("\n")}`,
      );
      return;
    }
    if (!isValidEmail(tempProfile.email)) {
      Alert.alert("Error", "Invalid email");
      return;
    }
    if (hasEmailChanged()) {
      if (!tempProfile.verificationPassword) {
        Alert.alert("Error", "Please enter your password to change the email.");
        return;
      }
      try {
        const response = await axios.put(
          `${data.API_URL}/users/me/email`,
          {
            email: tempProfile.email,
            password: tempProfile.verificationPassword,
          },
          { headers },
        );
      } catch (error) {
        Alert.alert("Error", "Invalid password.");
        return;
      }
    }

    if (
      kind === "password" &&
      tempProfile.newPassword !== "" &&
      tempProfile.previousPassword !== ""
    ) {
      try {
        const passwordResponse = await axios.put(
          `${data.API_URL}/users/me/password`,
          {
            new_password: tempProfile.newPassword,
            previous_password: tempProfile.previousPassword,
          },
          { headers },
        );
      } catch (error) {
        Alert.alert("Error", "Invalid previous password.");
        return;
      }
    }

    const profileResponse = await axios.put(
      `${data.API_URL}/users/me`,
      {
        username: tempProfile.username,
        first_name: tempProfile.firstName,
        last_name: tempProfile.lastName,
      },
      { headers },
    );
    setFirstName(tempProfile.firstName);
    setLastName(tempProfile.lastName);
    setEmail(tempProfile.email);
    setUsername(tempProfile.username);
    setNewPassword("");
    setPreviousPassword("");
    setVerificationPassword("");
    setIsEditProfileModalVisible(false);
  }

  const hasEmailChanged = () => {
    return tempProfile.email !== email;
  };

  const renderEditProfileModal = () => {
    return (
      <Modal
        animationType="slide"
        transparent={false}
        visible={isEditProfileModalVisible}
        onRequestClose={closeEditProfileModalWithoutSaving}
      >
        <View style={styles.modalContainer}>
          <View style={styles.editButtonsContainer}>
            <TouchableOpacity onPress={saveChanges}>
              <Text style={styles.saveText}>Save</Text>
            </TouchableOpacity>
            <TouchableOpacity onPress={closeEditProfileModalWithoutSaving}>
              <Text style={styles.closeText}>Close</Text>
            </TouchableOpacity>
          </View>
          <Text style={styles.editProfileHeader}>Edit Profile</Text>
          <View style={styles.inputContainer}>
            <Text style={styles.inputLabel}>Username</Text>
            <TextInput
              style={styles.input}
              value={tempProfile.username.toString()}
              onChangeText={(value) =>
                setTempProfile((prev) => ({ ...prev, username: value }))
              }
            />
          </View>
          {kind === "password" && (
            <View style={styles.inputContainer}>
              <Text style={styles.inputLabel}>Email</Text>
              <TextInput
                style={styles.input}
                value={tempProfile.email.toString()}
                onChangeText={(value) =>
                  setTempProfile((prev) => ({ ...prev, email: value }))
                }
                keyboardType="email-address"
              />
            </View>
          )}
          {hasEmailChanged() && (
            <View style={styles.inputContainer}>
              <Text style={styles.inputLabel}>
                Confirm password for email change
              </Text>
              <TextInput
                style={styles.input}
                value={tempProfile.verificationPassword}
                onChangeText={(value) =>
                  setTempProfile((prev) => ({
                    ...prev,
                    verificationPassword: value,
                  }))
                }
                secureTextEntry={true}
                placeholder="Enter your password"
              />
            </View>
          )}
          <View style={styles.inputContainer}>
            <Text style={styles.inputLabel}>First Name</Text>
            <TextInput
              style={styles.input}
              value={tempProfile.firstName.toString()}
              onChangeText={(value) =>
                setTempProfile((prev) => ({ ...prev, firstName: value }))
              }
            />
          </View>
          <View style={styles.inputContainer}>
            <Text style={styles.inputLabel}>Last Name</Text>
            <TextInput
              style={styles.input}
              value={tempProfile.lastName.toString()}
              onChangeText={(value) =>
                setTempProfile((prev) => ({ ...prev, lastName: value }))
              }
            />
          </View>
          {kind === "password" && (
            <View style={styles.inputContainer}>
              <Text style={styles.inputLabel}>New Password</Text>
              <TextInput
                style={styles.input}
                value={tempProfile.newPassword?.toString() ?? ""}
                onChangeText={(value) =>
                  setTempProfile((prev) => ({ ...prev, newPassword: value }))
                }
                secureTextEntry={true}
                placeholder="Enter your new password"
              />
            </View>
          )}
          {kind === "password" && (
            <View style={styles.inputContainer}>
              <Text style={styles.inputLabel}>Previous Password</Text>
              <TextInput
                style={styles.input}
                value={tempProfile.previousPassword?.toString() ?? ""}
                onChangeText={(value) =>
                  setTempProfile((prev) => ({
                    ...prev,
                    previousPassword: value,
                  }))
                }
                secureTextEntry={true}
                placeholder="Enter your previous password"
              />
            </View>
          )}
        </View>
      </Modal>
    );
  };

  const settingsModal = () => {
    return (
      <HeaderButtonComponent
        onPressAction={() => {
          setIsBottomModalVisible(true);
        }}
        logo={() => {
          return <Ionicons name="ios-settings" color={"white"} size={20} />;
        }}
        backgroundColor={"#8F5495"}
      />
    );
  };

  return (
    <LinearGradient
      colors={["#D0CBED", "#837AB1"]}
      style={{ ...styles.container, backgroundColor: "#D0CBED" }}
    >
      <HeaderComponent title={"Profil"} button={settingsModal()} />
      <View style={styles.detailsCard}>
        <Text style={styles.nameText}>
          {firstName} {lastName}
        </Text>
        <Text style={styles.nameText}>{username}</Text>
        <Text style={styles.detailsText}>
          <Ionicons name="mail" size={18} color="grey" /> {email}
        </Text>
        <Text style={styles.detailsText}>
          <Ionicons name="map" size={18} color="grey" /> Total Areas:{" "}
          {areaCount}
        </Text>
      </View>
      <Modal
        animationType="slide"
        transparent={true}
        visible={isBottomModalVisible}
        onRequestClose={() => {
          setIsBottomModalVisible(false);
        }}
      >
        <TouchableOpacity
          style={styles.bottomModalContainer}
          onPress={() => setIsBottomModalVisible(false)}
        >
          <TouchableOpacity
            activeOpacity={1}
            style={styles.modalView}
            onPress={(e) => e.stopPropagation()}
          >
            <Text style={styles.modalText}>Settings</Text>
            <MyButton
              name={"Sign Out"}
              ButtonStyle={styles.modalButton}
              TextStyle={styles.SignOutText}
              onPress={Disconnection}
            />
            <MyButton
              name={"Link services"}
              ButtonStyle={styles.modalButton}
              TextStyle={styles.linkText}
              onPress={navigateToAskService}
            />
            <MyButton
              name={"Delete Account"}
              ButtonStyle={styles.DeleteAccountText}
              TextStyle={styles.DeleteAccountText}
              onPress={DeleteAccount}
            />
            <MyButton
              name={"Annuler"}
              ButtonStyle={styles.modalCancelButton}
              TextStyle={styles.cancelText}
              onPress={() => setIsBottomModalVisible(false)}
            />
          </TouchableOpacity>
        </TouchableOpacity>
      </Modal>
      {renderEditProfileModal()}
      <MyButton
        name={"Edit profile"}
        ButtonStyle={styles.editButton}
        TextStyle={styles.editText}
        onPress={openEditProfileModal}
      />
    </LinearGradient>
  );
}
