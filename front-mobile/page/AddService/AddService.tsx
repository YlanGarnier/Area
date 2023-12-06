import React, { useState } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  Modal,
  TextInput,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  Alert,
} from "react-native";
import styles from "./style";
import * as services from "front-mobile/services.json";
import axios from "axios";
import * as data from "front-mobile/global.json";
import * as SecureStore from "expo-secure-store";
import { HeaderComponent } from "../../components/Header/HeaderComponent";
import { LinearGradient } from "expo-linear-gradient";
import AppModal from "../../components/AppModal/AppModal";
import LogoGithub from "../../assets/logo-github";
import LogoFacebook from "../../assets/logo-facebook";
import LogoDiscord from "../../assets/logo-discord";
import LogoSpotify from "../../assets/logo-spotify";
import LogoMiro from "../../assets/logo-miro";
import LogoTwitter from "../../assets/logo-twitter";
import LogoTwitch from "../../assets/logo-twitch";
import LogoNotion from "../../assets/logo-notion";
import LogoDropbox from "../../assets/logo-dropbox";
import LogoLinkedin from "../../assets/logo-linkedin";
import { HeaderButtonComponent } from "../../components/Header/HeaderButtonComponent";
import { Ionicons } from "@expo/vector-icons";
import LogoHttp from "../../assets/logo-http";
import LogoEthereum from "../../assets/logo-ethereum";
import LogoGmail from "../../assets/logo-gmail";

interface AppImage {
  name: string;
  component: JSX.Element;
}

const AddServicePage: React.FC = () => {
  const [selectedApp1, setSelectedApp1] = useState<string | null>("http");
  const [selectedApp2, setSelectedApp2] = useState<string | null>("github");
  const [isApp1ModalVisible, setIsApp1ModalVisible] = useState<boolean>(false);
  const [isApp2ModalVisible, setIsApp2ModalVisible] = useState<boolean>(false);
  const [selectedAction1, setSelectedAction1] =
    useState<string>("Select Action");
  const [isActionModalVisible, setIsActionModalVisible] =
    useState<boolean>(false);
  const [selectedReaction, setSelectedReaction] =
    useState<string>("Select Reaction");
  const [isReactionModalVisible, setIsReactionModalVisible] =
    useState<boolean>(false);
  const [isCreateModalVisible, setIsCreateModalVisible] =
    useState<boolean>(false);
  const [actionsForSelectedApp, setActionsForSelectedApp] = useState<any[]>([]);
  const [reactionsForSelectedApp, setReactionsForSelectedApp] = useState<any[]>(
    [],
  );
  const [inputParams, setInputParams] = useState<{ [key: string]: string }>({});
  const [inputParams2, setInputParams2] = useState<{ [key: string]: string }>(
    {},
  );
  const [areaName, setAreaName] = useState<string>("");
  const [isInputFocused, setInputFocused] = useState(false);

  const handleInputFocus = () => {
    setInputFocused(true);
  };

  const handleInputBlur = () => {
    setInputFocused(false);
  };

  const appImages: AppImage[] = [
    { name: "http", component: <LogoHttp width={60} height={60} /> },
    { name: "github", component: <LogoGithub width={60} height={60} /> },
    { name: "ethereum", component: <LogoEthereum width={60} height={60} /> },
    { name: "google_gmail", component: <LogoGmail width={60} height={60} /> },
    { name: "discord", component: <LogoDiscord width={60} height={60} /> },
    { name: "facebook", component: <LogoFacebook width={60} height={60} /> },
    { name: "spotify", component: <LogoSpotify width={60} height={60} /> },
    { name: "miro", component: <LogoMiro width={60} height={60} /> },
    { name: "twitch", component: <LogoTwitch width={60} height={60} /> },
    { name: "notion", component: <LogoNotion width={60} height={60} /> },
    { name: "dropbox", component: <LogoDropbox width={60} height={60} /> },
    { name: "linkedin", component: <LogoLinkedin width={60} height={60} /> },
    { name: "twitter", component: <LogoTwitter width={60} height={60} /> },
  ];

  const handleActionSelection = (actionName: string) => {
    const selectedAction = actionsForSelectedApp.find(
      (route) => route.name === actionName,
    );
    if (selectedAction && selectedAction.params) {
      const emptyParams = Object.fromEntries(
        Object.keys(selectedAction.params).map((key) => [key, ""]),
      );
      setInputParams(emptyParams);
    } else {
      setInputParams({});
    }
    setSelectedAction1(actionName);
    setIsActionModalVisible(false);
  };

  const handleReactionSelection = (reactionName: string) => {
    const selectedReaction = reactionsForSelectedApp.find(
      (route) => route.name === reactionName,
    );
    if (selectedReaction && selectedReaction.target) {
      setInputParams2({ [selectedReaction.target]: "" });
    } else {
      setInputParams2({});
    }
    setSelectedReaction(reactionName);
    setIsReactionModalVisible(false);
  };

  const handleApp1Selection = (appName: string) => {
    const appActions = services.actions.find(
      (action) => action.name === appName,
    );
    if (appActions && appActions.routes.length > 0) {
      setActionsForSelectedApp(appActions.routes);
    } else {
      setActionsForSelectedApp([]);
    }
    setSelectedApp1(appName);
    setSelectedAction1("Choose action");
    setIsApp1ModalVisible(false);
    setInputParams({});
  };

  const handleApp2Selection = (appName: string) => {
    const appReactions = services.reactions.find(
      (reaction) => reaction.name === appName,
    );
    if (appReactions && appReactions.routes.length > 0) {
      setReactionsForSelectedApp(appReactions.routes);
    } else {
      setReactionsForSelectedApp([]);
    }
    setSelectedReaction("Select Reaction");
    setInputParams2({});
    setSelectedApp2(appName);
    setIsApp2ModalVisible(false);
  };

  const renderReactionInputs = () => {
    const paramKey = Object.keys(inputParams2)[0];
    if (paramKey) {
      return (
        <View>
          <TextInput
            style={styles.inputArea}
            placeholder={`Enter ${paramKey}`}
            onChangeText={(value) => setInputParams2({ [paramKey]: value })}
            value={inputParams2[paramKey] || ""}
            onFocus={handleInputFocus}
            onBlur={handleInputBlur}
          />
        </View>
      );
    }
    return null;
  };

  const isValidForm = (): boolean => {
    if (
      !areaName ||
      !selectedApp1 ||
      !selectedApp2 ||
      selectedAction1 === "Select Action" ||
      selectedAction1 === "Choose action" ||
      selectedReaction === "Select Reaction" ||
      selectedReaction === "Choose Reaction"
    ) {
      return false;
    }
    for (const key in inputParams) {
      if (inputParams[key] === "") {
        return false;
      }
    }
    for (const key in inputParams2) {
      if (inputParams2[key] === "") {
        return false;
      }
    }

    return true;
  };

  const resetState = () => {
    setAreaName("");
    setSelectedApp1("http");
    setSelectedApp2("github");
    setSelectedAction1("Select Action");
    setSelectedReaction("Select Reaction");
    setInputParams({});
    setInputParams2({});
  };

  const handleCreate = async () => {
    const headers = {
      Authorization: `Bearer ${await SecureStore.getItemAsync("jwtToken")}`,
    };
    const actionToken = await SecureStore.getItemAsync(`${selectedApp1}Token`);
    const reactionToken = await SecureStore.getItemAsync(
      `${selectedApp2}Token`,
    );
    if (
      (!actionToken || !reactionToken) &&
      selectedApp1 !== "http" &&
      selectedApp2 !== "http"
    ) {
      Alert.alert(
        `Impossible to create area: "${selectedApp1}" or "${selectedApp2}" account is not linked`,
      );
      return;
    }

    if (isValidForm()) {
      const actionService = selectedApp1;
      const reactionService = selectedApp2;
      const actionRoute = selectedAction1;
      const reactionRoute = selectedReaction;

      const payload = {
        name: areaName,
        action: {
          service: actionService,
          route: actionRoute,
          params: { ...inputParams },
        },
        reaction: {
          service: reactionService,
          route: reactionRoute,
          target: Object.values(inputParams2)[0] || "",
        },
      };
      try {
        const response = await axios.post(`${data.API_URL}/area/new`, payload, {
          headers: headers,
        });
        resetState();
      } catch (error) {
        Alert.alert("Error posting data:");
        resetState();
      }

      setSelectedApp1("http");
      setSelectedApp2("github");
      setIsCreateModalVisible(false);
    }
  };

  const closeModal = (setter: (value: boolean) => void) => {
    setter(false);
  };

  const selectedAppComponent1 = appImages.find(
    (img) => img.name === selectedApp1,
  )?.component;
  const selectedAppComponent2 = appImages.find(
    (img) => img.name === selectedApp2,
  )?.component;

  const resetButtonModal = () => {
    return (
      <HeaderButtonComponent
        onPressAction={resetState}
        logo={() => {
          return <Ionicons name="ios-refresh" color={"white"} size={20} />;
        }}
        backgroundColor={"#8F5495"}
      />
    );
  };

  const getAvailableActionImages = () => {
    return appImages.filter((image) => {
      const action = services.actions.find((a) => a.name === image.name);
      return action && action.routes.length > 0;
    });
  };

  const getAvailableReactionImages = () => {
    return appImages.filter((image) => {
      const reaction = services.reactions.find((r) => r.name === image.name);
      return reaction && reaction.routes.length > 0;
    });
  };

  const availableActions = getAvailableActionImages();
  const availableReactions = getAvailableReactionImages();

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      style={{ flex: 1, height: 1000 }}
    >
      <LinearGradient colors={["#D0CBED", "#837AB1"]} style={styles.container}>
        <HeaderComponent title={"Add area"} button={resetButtonModal()} />
        <Text style={styles.sectionTitle}>Area Name</Text>
        <View>
          <TextInput
            style={styles.inputArea}
            placeholder="Enter Area Name"
            value={areaName}
            onChangeText={(text) => setAreaName(text)}
            onFocus={handleInputFocus}
            onBlur={handleInputBlur}
          />
        </View>
        <Text style={styles.sectionTitle}>Action</Text>
        <ScrollView style={styles.scrollViewStyle}>
          <View style={styles.actionSection}>
            <TouchableOpacity
              style={styles.appSelector}
              onPress={() => setIsApp1ModalVisible(true)}
            >
              {selectedAppComponent1}
            </TouchableOpacity>

            <AppModal
              isVisible={isApp1ModalVisible}
              onClose={() => setIsApp1ModalVisible(false)}
              appImages={availableActions}
              onSelectApp={(appId) => handleApp1Selection(appId)}
            />

            <Modal
              visible={isActionModalVisible}
              animationType="slide"
              transparent
              onRequestClose={() => closeModal(setIsActionModalVisible)}
            >
              <View style={styles.modalOverlay}>
                <View style={styles.dropdownModal}>
                  {actionsForSelectedApp.map((route) => (
                    <TouchableOpacity
                      key={route.name}
                      onPress={() => handleActionSelection(route.name)}
                      style={
                        selectedAction1 === route.name
                          ? styles.selectedItem
                          : null
                      }
                    >
                      <View style={styles.modalItemContent}>
                        <Text style={styles.modalItemText}>{route.name}</Text>
                      </View>
                    </TouchableOpacity>
                  ))}
                  <TouchableOpacity
                    onPress={() => closeModal(setIsActionModalVisible)}
                  >
                    <Text style={styles.modalItemText}>Close</Text>
                  </TouchableOpacity>
                </View>
              </View>
            </Modal>

            <TouchableOpacity onPress={() => setIsActionModalVisible(true)}>
              <Text style={styles.dropdownText}>
                {selectedAction1 === "Choose action"
                  ? "Choose action"
                  : selectedAction1}
              </Text>
            </TouchableOpacity>

            {Object.keys(inputParams).length > 0 && (
              <View>
                {Object.keys(inputParams).map((key, index) => (
                  <TextInput
                    key={index}
                    style={styles.inputArea}
                    placeholder={`Enter ${key}`}
                    onChangeText={(value) =>
                      setInputParams({ ...inputParams, [key]: value })
                    }
                    value={inputParams[key] || ""}
                    onFocus={handleInputFocus}
                    onBlur={handleInputBlur}
                  />
                ))}
              </View>
            )}
          </View>
        </ScrollView>

        <View style={styles.separator} />
        <ScrollView>
          <Text style={styles.sectionTitle}>Reaction</Text>
          <View style={styles.reactionSection}>
            <TouchableOpacity
              style={styles.appSelector}
              onPress={() => setIsApp2ModalVisible(true)}
            >
              {selectedAppComponent2}
            </TouchableOpacity>
            <AppModal
              isVisible={isApp2ModalVisible}
              onClose={() => setIsApp2ModalVisible(false)}
              appImages={availableReactions}
              onSelectApp={(appId) => handleApp2Selection(appId)}
            />
            <Modal
              visible={isReactionModalVisible}
              animationType="slide"
              transparent
              onRequestClose={() => closeModal(setIsReactionModalVisible)}
            >
              <View style={styles.modalOverlay}>
                <View style={styles.dropdownModal}>
                  {reactionsForSelectedApp.map((route) => (
                    <TouchableOpacity
                      key={route.name}
                      onPress={() => handleReactionSelection(route.name)}
                      style={
                        selectedReaction === route.name
                          ? styles.selectedItem
                          : null
                      }
                    >
                      <View style={styles.modalItemContent}>
                        <Text style={styles.modalItemText}>{route.name}</Text>
                      </View>
                    </TouchableOpacity>
                  ))}
                  <TouchableOpacity
                    onPress={() => closeModal(setIsReactionModalVisible)}
                  >
                    <Text style={styles.modalItemText}>Close</Text>
                  </TouchableOpacity>
                </View>
              </View>
            </Modal>
            <TouchableOpacity onPress={() => setIsReactionModalVisible(true)}>
              <Text style={styles.dropdownText}>
                {selectedReaction === "Select Reaction"
                  ? "Choose Reaction"
                  : selectedReaction}
              </Text>
            </TouchableOpacity>
            {selectedApp2 && renderReactionInputs()}
          </View>
        </ScrollView>
        {(!isInputFocused || isValidForm()) && (
          <TouchableOpacity
            style={[styles.button, !isValidForm() && styles.buttonDisabled]}
            onPress={() => handleCreate()}
            disabled={!isValidForm()}
          >
            <Text style={styles.buttonText}>Create</Text>
          </TouchableOpacity>
        )}
      </LinearGradient>
    </KeyboardAvoidingView>
  );
};

export default AddServicePage;
