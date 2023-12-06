import React, { useState } from "react";
import { Alert, Text, TouchableOpacity, View } from "react-native";
import styles from "./style";
import { StackNavigationProp } from "@react-navigation/stack";
import { RouteProp } from "@react-navigation/native";
import { RootStackParamList } from "../../App";
import { LinearGradient } from "expo-linear-gradient";
import { HeaderComponent } from "../../components/Header/HeaderComponent";
import { HeaderButtonComponent } from "../../components/Header/HeaderButtonComponent";
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
import LogoHttp from "../../assets/logo-http";
import LogoGmail from "../../assets/logo-gmail";
import * as SecureStore from "expo-secure-store";
import * as data from "front-mobile/global.json";
import axios from "axios";
import DeleteConfirmationModal from "../../components/confirmModal/confirmModal";

export interface AreaPageRouteParams {
  paramName: { name: string };
  id: string;
  actionName: string;
  reactionName: string;
  route_action_service: string;
  route_reaction_service: string;
}

type AreaNavigationProp = StackNavigationProp<RootStackParamList, "Area">;
type AreaRouteProp = RouteProp<RootStackParamList, "Area">;

type AreaPageProps = {
  navigation: AreaNavigationProp;
  route: AreaRouteProp;
};

function AreaPage({ route, navigation }: AreaPageProps) {
  const {
    paramName,
    id,
    actionName,
    reactionName,
    route_action_service,
    route_reaction_service,
  } = route.params;

  const appLogos = {
    http: <LogoHttp width={100} height={100} />,
    google_gmail: <LogoGmail width={100} height={100} />,
    github: <LogoGithub width={100} height={100} />,
    discord: <LogoDiscord width={100} height={100} />,
    facebook: <LogoFacebook width={100} height={100} />,
    twitch: <LogoTwitch width={100} height={100} />,
    miro: <LogoMiro width={100} height={100} />,
    spotify: <LogoSpotify width={100} height={100} />,
    notion: <LogoNotion width={100} height={100} />,
    dropbox: <LogoDropbox width={100} height={100} />,
    linkedin: <LogoLinkedin width={100} height={100} />,
    twitter: <LogoTwitter width={100} height={100} />,
  };

  function findLogo(logoName) {
    return appLogos[logoName.toLowerCase()] || null;
  }

  const actionLogo = findLogo(actionName);
  const reactionLogo = findLogo(reactionName);

  const goingBackToLobby = () => {
    return (
      <HeaderButtonComponent
        onPressAction={() => {
          navigation.navigate("Lobby");
        }}
        logo={() => {
          return <Ionicons name="arrow-back" color={"white"} size={20} />;
        }}
        backgroundColor={"#8F5495"}
      />
    );
  };

  function capitalizeFirstLetter(text: string) {
    return text.charAt(0).toUpperCase() + text.slice(1);
  }

  const renderAppZone = (
    appLogo,
    appAction: string,
    route: string,
    str: string,
  ) => {
    return (
      <View style={styles.appZone}>
        <View>{appLogo}</View>
        <Text style={styles.appAction}>{capitalizeFirstLetter(appAction)}</Text>
        <Text style={styles.text}>
          {str} {route}
        </Text>
      </View>
    );
  };

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
      navigation.navigate("Lobby");
    }
    setDeleteConfirmationVisible(false);
  };

  return (
    <LinearGradient colors={["#D0CBED", "#837AB1"]} style={styles.container}>
      <HeaderComponent title={paramName.name} button={goingBackToLobby()} />
      <View style={styles.contentContainer}>
        {renderAppZone(
          actionLogo,
          actionName,
          route_action_service,
          "Action :",
        )}
        <View style={styles.verticalLine} />
        {renderAppZone(
          reactionLogo,
          reactionName,
          route_reaction_service,
          "Reaction :",
        )}
      </View>
      <TouchableOpacity style={styles.deleteButton} onPress={handleDelete}>
        <Ionicons name="trash-bin" size={50} color="#725A8A" />
      </TouchableOpacity>
      <DeleteConfirmationModal
        isVisible={isDeleteConfirmationVisible}
        onClose={() => setDeleteConfirmationVisible(false)}
        onDelete={confirmDelete}
      />
    </LinearGradient>
  );
}

export default AreaPage;
