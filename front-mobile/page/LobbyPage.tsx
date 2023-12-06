import { Text, View } from "react-native";
import * as React from "react";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import HomePage from "./HomePage/home";
import { Ionicons } from "@expo/vector-icons";
import { StackScreenProps } from "@react-navigation/stack";
import AddService from "./AddService/AddService";
import ProfilePage from "./ProfilePage/ProfilePage";
import { useEffect } from "react";
import * as SecureStore from "expo-secure-store";
import {
  fetchAndStoreAllServiceTokens,
  initToken,
} from "../Functions/tokenUtils";

export type RootTabParamList = {
  Services: undefined;
  "Add Area": undefined;
  Profile: undefined;
};

const Tab = createBottomTabNavigator<RootTabParamList>();

function getIconName(routeName: string, focused: boolean): string {
  switch (routeName) {
    case "Services":
      return focused ? "home" : "home-outline";
    case "Add Area":
      return focused ? "add" : "add-outline";
    case "Profile":
      return focused ? "person" : "person-outline";
    default:
      throw new Error(`Invalid route name: ${routeName}`);
  }
}

export default function LobbyPage({
  navigation,
  route,
}: StackScreenProps<any, "Lobby">) {
  const initialTab = route.params?.initialTab || "Services";

  useEffect(() => {
    if (SecureStore.getItemAsync("jwtToken") !== null) {
      fetchAndStoreAllServiceTokens();
    }
  }, []);
  return (
    <Tab.Navigator
      initialRouteName={initialTab}
      screenOptions={({ route }) => ({
        tabBarIcon: ({ focused, color, size }) => {
          const iconName = getIconName(route.name, focused);
          return <Ionicons name={iconName as any} size={size} color={color} />;
        },
        tabBarActiveTintColor: "#8F5495",
        tabBarInactiveTintColor: "#b9cbd0",
        headerShown: false,
        tabBarShowLabel: false,
        tabBarStyle: {
          shadowOffset: {
            width: 0,
            height: 12,
          },
          shadowOpacity: 0.58,
          shadowRadius: 16.0,
          elevation: 24,
          borderTopLeftRadius: 21,
          borderTopRightRadius: 21,
          backgroundColor: "white",
          position: "absolute",
          bottom: 0,
          padding: 10,
          width: "100%",
          height: 60,
          zIndex: 0,
        },
      })}
    >
      <Tab.Screen
        name={"Services"}
        component={HomePage}
        options={{ headerShown: false }}
      />
      <Tab.Screen
        name={"Add Area"}
        component={AddService}
        options={{ headerShown: false }}
      />
      <Tab.Screen
        name={"Profile"}
        component={ProfilePage}
        options={{ headerShown: false }}
      />
    </Tab.Navigator>
  );
}
