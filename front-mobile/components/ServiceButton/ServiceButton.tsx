import React, { useState, useEffect } from "react";
import { TouchableOpacity, Text, Alert, Linking } from 'react-native';
import * as SecureStore from 'expo-secure-store';
import axios from "axios";
import * as data from "front-mobile/global.json";

interface ServiceButtonProps {
    onPressLink: string;
    logo: any;
    text: string;
    tokenKey: string;
    style?: object;
    textStyle?: object;
    isLinked?: boolean;
}

const ServiceButton: React.FC<ServiceButtonProps> = ({ onPressLink, logo, text, tokenKey, style, textStyle, isLinked }) => {
    const Color = isLinked ? '#8F5495' : '#C3A6E5';

    const handleURL = (event: { url: string }) => {
        SecureStore.setItemAsync(tokenKey, "set");
    };

    useEffect(() => {
        const subscription = Linking.addEventListener('url', handleURL);
        return () => {
            subscription.remove();
        };
    }, []);

    const Redirected = async () => {
        if (isLinked) {
            Alert.alert(
                "Unlink Account",
                "Do you want to unlink this account?",
                [
                    {
                        text: "Cancel",
                        style: "cancel"
                    },
                    {
                        text: "Unlink",
                        onPress: async () => {
                            const headers = {
                                Authorization: `Bearer ${await SecureStore.getItemAsync('jwtToken')}`
                            };
                            await SecureStore.deleteItemAsync(tokenKey);
                            await axios.delete(
                                `${data.API_URL}/users/me/services/${tokenKey.slice(0, -5)}`,
                                { headers }
                            );
                        }
                    }
                ],
                { cancelable: false }
            );
        } else {
            await Linking.openURL(`${onPressLink}`);
        }
    };

    return (
        <TouchableOpacity style={style} onPress={Redirected}>
            {logo}
            <Text style={[textStyle, { color: Color }]}>
                {isLinked ? 'Linked' : text}
            </Text>
        </TouchableOpacity>
    );
};

export default ServiceButton;
