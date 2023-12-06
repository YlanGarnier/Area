import { ScrollView, Text, Alert } from "react-native";
import { useCallback, useEffect, useState } from "react";
import styles from "./style";
import ServiceCompo from "../../components/ServiceCompo/ServiceCompo";
import {
  NavigationProp,
  useFocusEffect,
  useNavigation,
} from "@react-navigation/native";
import { StackScreenProps } from "@react-navigation/stack";
import * as SecureStore from "expo-secure-store";
import axios from "axios";
import * as data from "front-mobile/global.json";
import { ServiceToggleProvider } from "../../components/Context/ToggleContext";
import { HeaderComponent } from "../../components/Header/HeaderComponent";
import { LinearGradient } from "expo-linear-gradient";
import { RootStackParamList } from "../../App";

export default function HomePage({ route }: StackScreenProps<any, "Services">) {
  const navigation = useNavigation<NavigationProp<RootStackParamList>>();
  const [services, setServices] = useState([]);
  const [servicesUpdatedAt, setServicesUpdatedAt] = useState(0);

  async function GetServices() {
    try {
      const headers = {
        Authorization: `Bearer ${await SecureStore.getItemAsync("jwtToken")}`,
      };
      const response = await axios.get(`${data.API_URL}/users/me/areas`, {
        headers,
      });
      if (response.data && Array.isArray(response.data)) {
        setServices(response.data);
      } else {
        Alert.alert("error from API");
      }
    } catch (error) {
      console.error(error);
    }
  }

  useFocusEffect(
    useCallback(() => {
      GetServices();
    }, []),
  );

  useEffect(() => {
    GetServices();
  }, [servicesUpdatedAt]);

  return (
    <LinearGradient
      colors={["#D0CBED", "#837AB1"]}
      style={{ ...styles.container, backgroundColor: "#D0CBED" }}
    >
      <HeaderComponent title={"Areas"} />
      <ScrollView style={styles.scrollView}>
        <ServiceToggleProvider>
          {services.length === 0 ? (
            <Text style={styles.noAreasText}>You don't have any areas.</Text>
          ) : (
            services.map((service, index) => (
              <ServiceCompo
                key={index}
                id={service.id}
                name={service.name}
                actionName={service.action_service}
                reactionName={service.reaction_service}
                route_action_service={service.route_action_service}
                route_reaction_service={service.route_reaction_service}
                navigation={navigation}
                route={route}
                onServiceDelete={() => {
                  setServicesUpdatedAt(Date.now());
                }}
              />
            ))
          )}
        </ServiceToggleProvider>
      </ScrollView>
    </LinearGradient>
  );
}
