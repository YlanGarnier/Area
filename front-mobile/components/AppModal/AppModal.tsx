import React from 'react';
import { View, Text, TouchableOpacity, Modal, ScrollView } from 'react-native';
import styles from './style';

interface AppModalProps {
    isVisible: boolean;
    onClose: () => void;
    appImages: { id: string; component: JSX.Element }[];
    onSelectApp: (appId: string) => void;
}

const AppModal: React.FC<AppModalProps> = ({ isVisible, onClose, appImages, onSelectApp }) => {
    const chunk = (array, size) => {
        const chunkedArr = [];
        let index = 0;
        while (index < array.length) {
            chunkedArr.push(array.slice(index, size + index));
            index += size;
        }
        return chunkedArr;
    };

    const iconRows = chunk(appImages, 4);

    return (
        <Modal
            visible={isVisible}
            animationType="slide"
            transparent
            onRequestClose={onClose}
        >
            <View style={styles.modalOverlay}>
                <View style={styles.dropdownModal}>
                    <Text>Select an App Image:</Text>
                    <ScrollView>
                        {iconRows.map((row, rowIndex) => (
                            <View key={rowIndex} style={styles.iconRow}>
                                {row.map((img) => (
                                    <TouchableOpacity key={img.name} style={styles.iconInRow} onPress={() => onSelectApp(img.name)}>
                                        {img.component}
                                    </TouchableOpacity>
                                ))}
                            </View>
                        ))}
                    </ScrollView>
                    <TouchableOpacity onPress={onClose}>
                        <Text>Close</Text>
                    </TouchableOpacity>
                </View>
            </View>
        </Modal>
    );
};

export default AppModal;
