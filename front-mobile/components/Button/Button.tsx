import {Text, TouchableOpacity, StyleProp, TextStyle, ViewStyle} from "react-native";
import React from "react";

type MyButtonProps = {
    name: string;
    ButtonStyle: StyleProp<ViewStyle>;
    TextStyle: StyleProp<TextStyle>;
    onPress: () => void;
};

export default function MyButton({name, ButtonStyle, TextStyle, onPress}: MyButtonProps) {
    return (
        <TouchableOpacity style={ButtonStyle} onPress={onPress}>
            <Text style={TextStyle}>{name}</Text>
        </TouchableOpacity>
    );
}
