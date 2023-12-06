import React, { useEffect, useState } from "react";
import { NavigationContainer, useNavigation } from "@react-navigation/native";
import { createStackNavigator } from "@react-navigation/stack";
import Register from "./page/RegisterPage/register";
import RegisterInfo from "./page/RegisterInfo/RegisterInfo";
import WelcomePage from "./page/WelcomePage/WelcomePage";
import LobbyPage from "./page/LobbyPage";
import AreaPage from "./page/AreaPage/area";
import ProfilePage from "./page/ProfilePage/ProfilePage";
import { ServiceToggleProvider } from "./components/Context/ToggleContext";
import { AreaPageRouteParams } from "./page/AreaPage/area";
import AskService from "./page/AskServicePage/AskServicePage";
import * as SecureStore from "expo-secure-store";
import {
  fetchAndStoreAllServiceTokens,
  initToken,
} from "./Functions/tokenUtils";

export type RootStackParamList = {
  Welcome: undefined;
  Register: undefined;
  RegisterInfo: undefined;
  AskService: { fromRegister?: boolean };
  Lobby: { initialTab?: "Services" | "Add Area" | "Profile" };
  Area: AreaPageRouteParams;
  Home: undefined;
  Services: undefined;
  "Add Area": undefined;
  Profile: undefined;
};

const Stack = createStackNavigator<RootStackParamList>();

function App() {
  useEffect(() => {
    if (SecureStore.getItemAsync("jwtToken") !== null) {
      initToken();
    }
  }, []);
  return (
    <ServiceToggleProvider>
      <NavigationContainer>
        <Stack.Navigator initialRouteName="Welcome">
          <Stack.Screen
            name="Welcome"
            component={WelcomePage}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name="Register"
            component={Register}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name="RegisterInfo"
            component={RegisterInfo}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name="AskService"
            component={AskService}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name="Lobby"
            component={LobbyPage}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name="Area"
            component={AreaPage}
            options={{ headerShown: false }}
          />
          <Stack.Screen name="Profile" component={ProfilePage} />
        </Stack.Navigator>
      </NavigationContainer>
    </ServiceToggleProvider>
  );
}

export default App;
