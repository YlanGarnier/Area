import React, { useState, useEffect } from 'react';
import { TouchableOpacity, Animated, View, Text, Image } from 'react-native';
import styles from "./style";

interface OnOffProps {
    isActive: boolean;
    onToggle: () => void;
}

const OnOff: React.FC<OnOffProps> = ({ isActive, onToggle }) => {
    const [localIsActive, setLocalIsActive] = useState(isActive);
    const sliderPosition = new Animated.Value(isActive ? 82 : 2);

    useEffect(() => {
        Animated.timing(sliderPosition, {
            toValue: localIsActive ? 82 : 2,
            duration: 200,
            useNativeDriver: false,
        }).start();
    }, [localIsActive]);


    const handleToggle = () => {
        setLocalIsActive(!localIsActive);
    }

    return (
        <View style={styles.buttonContainer}>
            <TouchableOpacity style={[styles.button, isActive ? styles.on : styles.off]} onPress={onToggle}>
                <Text style={styles.label}>{isActive ? "On" : "Off"}</Text>
                <Animated.View
                    style={[
                        styles.slider,
                        { left: sliderPosition },
                        isActive ? styles.sliderOn : styles.sliderOff
                    ]}
                ></Animated.View>
            </TouchableOpacity>
        </View>
    );
}

export default OnOff;
