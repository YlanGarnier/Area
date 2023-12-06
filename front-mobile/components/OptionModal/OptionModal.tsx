import React from 'react';
import { View, Modal, FlatList, TouchableOpacity, Text } from "react-native";
import styles from "./style";

interface OptionModalProps {
    visible: boolean;
    options: string[];
    onOptionSelected: (option: string) => void;
    closeModal: () => void;
}

const OptionModal: React.FC<OptionModalProps> = ({ visible, options, onOptionSelected, closeModal }) => {
    return (
        <Modal animationType="slide" transparent={true} visible={visible}>
            <View style={styles.centeredView}>
                <View style={styles.modalView}>
                    <FlatList
                        data={options}
                        keyExtractor={(item) => item}
                        renderItem={({ item }) => (
                            <TouchableOpacity
                                style={styles.modalOption}
                                onPress={() => {
                                    onOptionSelected(item);
                                    closeModal();
                                }}
                            >
                                <Text style={styles.modalText}>{item}</Text>
                            </TouchableOpacity>
                        )}
                    />
                </View>
            </View>
        </Modal>
    );
}

export default OptionModal;
