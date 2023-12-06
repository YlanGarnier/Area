import React, { useEffect, useState } from "react";
import { Text, View, TouchableOpacity, Alert } from "react-native";
import styles from "./style";
import { useServiceToggle } from "../Context/ToggleContext";
import { Ionicons } from "@expo/vector-icons";
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
import * as SecureStore from "expo-secure-store";
import * as data from "front-mobile/global.json";
import axios from "axios";
import DeleteConfirmationModal from "../confirmModal/confirmModal";

interface ServiceCompoProps {
  id: string;
  name: string;
  actionName: string;
  reactionName: string;
  route_action_service: string;
  route_reaction_service: string;
  navigation: any;
  route: any;
  onServiceDelete: () => void;
}

export default function ServiceCompo({
  id,
  name,
  actionName,
  reactionName,
  route_action_service,
  route_reaction_service,
  navigation,
  route,
  onServiceDelete,
}: ServiceCompoProps) {
  const appLogos = [
    { name: "google_gmail", logo: <LogoGmail width={40} height={40} /> },
    { name: "github", logo: <LogoGithub width={40} height={40} /> },
    { name: "discord", logo: <LogoDiscord width={40} height={40} /> },
    { name: "facebook", logo: <LogoFacebook width={40} height={40} /> },
    { name: "twitch", logo: <LogoTwitch width={40} height={40} /> },
    { name: "miro", logo: <LogoMiro width={40} height={40} /> },
    { name: "spotify", logo: <LogoSpotify width={40} height={40} /> },
    { name: "notion", logo: <LogoNotion width={40} height={40} /> },
    { name: "dropbox", logo: <LogoDropbox width={40} height={40} /> },
    { name: "linkedin", logo: <LogoLinkedin width={40} height={40} /> },
    { name: "twitter", logo: <LogoTwitter width={40} height={40} /> },
  ];

  function findLogo(logoName) {
    const lowerCaseName = logoName.toLowerCase();
    const appLogo = appLogos.find((app) => app.name === lowerCaseName);
    return appLogo ? appLogo.logo : null;
  }

  const actionLogo = findLogo(actionName);
  const reactionLogo = findLogo(reactionName);
  const { toggles, setToggle } = useServiceToggle();

  const currentToggleState = toggles[name] || false;

  useEffect(() => {
    const unsubscribe = navigation.addListener("focus", () => {
      if (route.params && route.params.toggleState !== undefined) {
        setToggle(name, route.params.toggleState);
      }
    });

    return () => {
      unsubscribe();
    };
  }, [navigation, route, name]);

  const [isDeleteConfirmationVisible, setDeleteConfirmationVisible] =
    useState(false);

  const handleDelete = async () => {
    setDeleteConfirmationVisible(true);
  };

  const confirmDelete = async () => {
    const headers = {
      Authorization: `Bearer ${await SecureStore.getItemAsync("jwtToken")}`,
    };
    const response = await axios.delete(`${data.API_URL}/area/${id}`, {
      headers,
    });
    if (response.status !== 204) {
      Alert.alert("Error", "An error occurred while deleting the area");
    } else {
      onServiceDelete();
    }
    setDeleteConfirmationVisible(false);
  };

  const handleServicePress = () => {
    navigation.navigate("Area", {
      paramName: { name },
      id,
      actionName,
      reactionName,
      route_action_service,
      route_reaction_service,
    });
  };

  return (
    <TouchableOpacity onPress={handleServicePress}>
      <View style={styles.container}>
        <View style={styles.content}>
          <View style={styles.titleContainer}>
            <View style={styles.leftContainer}>
              {actionLogo}
              {reactionLogo}
            </View>
            <Text style={styles.title} ellipsizeMode="tail" numberOfLines={1}>
              {name}
            </Text>
          </View>
          <TouchableOpacity style={styles.deleteButton} onPress={handleDelete}>
            <Ionicons name="trash-bin" size={30} color="#725A8A" />
          </TouchableOpacity>
        </View>
      </View>
      <DeleteConfirmationModal
        isVisible={isDeleteConfirmationVisible}
        onClose={() => setDeleteConfirmationVisible(false)}
        onDelete={confirmDelete}
      />
    </TouchableOpacity>
  );
}
