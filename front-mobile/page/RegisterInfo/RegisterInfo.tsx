import { useState, useEffect } from "react";
import { KeyboardAvoidingView, Platform, Image, Text, TextInput, TouchableOpacity, View, Button } from "react-native";
import { StatusBar } from "expo-status-bar";
import styles from './style'
import axios from "axios";
import { StackScreenProps } from '@react-navigation/stack';
import * as SecureStore from "expo-secure-store";
import * as data from "front-mobile/global.json";

type RootStackParamList = {
    Welcome: undefined;
    Register: undefined;
    RegisterInfo: undefined;
    AskService: undefined
    Lobby: undefined;
    Area: undefined;
    Home: undefined;
};

export default function RegisterInfo({ navigation, route }: StackScreenProps<RootStackParamList, 'RegisterInfo'>) {
    const [username, setUsername] = useState("");
    const [firstName, setFirstname] = useState("");
    const [lastName, setLastname] = useState("");
    const [message, setMessage] = useState("");


    const onSubmitFormHandler = async () => {
        try {
            const headers = {"Authorization" : `Bearer ${await SecureStore.getItemAsync('jwtToken')}`}
            const response = await axios.put(`${data.API_URL}/users/me`, { first_name:firstName, last_name:lastName, username:username }, {headers});
            if (!response.data?.status) {
                navigation.reset({
                    index: 0,
                    routes: [{ name: "AskService", params: { fromRegister: true } }],
                });
            } else {
                setMessage("error invalid credentials");
            }
        } catch {
            setMessage("error on request");
        }
    };

    return (
        <KeyboardAvoidingView style={styles.container} behavior={Platform.OS === "ios" ? "padding" : "height"} enabled>
            <View style={styles.container}>
                <Image style={styles.background} source={require('front-mobile/assets/HomeBg.jpeg')} />
                <Text style={styles.catchphrase}>We're almost done!</Text>
                <View style={styles.middle}>
                    <View style={styles.inputContainer}>
                        <Text style={styles.label}>First Name</Text>
                        <TextInput onChangeText={setFirstname} style={styles.input} />
                    </View>
                    <View style={styles.inputContainer}>
                        <Text style={styles.label}>Last Name</Text>
                        <TextInput onChangeText={setLastname} style={styles.input} />
                    </View>
                    <View style={styles.inputContainer}>
                        <Text style={styles.label}>Username</Text>
                        <TextInput onChangeText={setUsername} style={styles.input} />
                    </View>
                    <View style={styles.messageContainer}>
                        <Text style={styles.errmessage}>
                            {message}
                        </Text>
                    </View>
                    <TouchableOpacity style={styles.signButton} onPress={onSubmitFormHandler}>
                        <Text style={styles.signinText}>Next</Text>
                    </TouchableOpacity>
                </View>
                <StatusBar style="auto" />
            </View>
        </KeyboardAvoidingView>
    );
}
