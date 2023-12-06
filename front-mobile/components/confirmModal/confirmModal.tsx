import React from "react";
import { Modal, View, Text, Button, TouchableOpacity } from "react-native";
import styles from "./style";

function DeleteConfirmationModal({ isVisible, onClose, onDelete }) {
  return (
    <Modal visible={isVisible} transparent animationType="slide">
      <View style={styles.container}>
        <View style={styles.modal}>
          <Text style={styles.text}>
            Are you sure you want to delete this area ?
          </Text>
          <TouchableOpacity onPress={onDelete}>
            <Text style={styles.confirmButton}>Delete</Text>
          </TouchableOpacity>
          <Button title="Cancel" onPress={onClose} color="#8F5495" />
        </View>
      </View>
    </Modal>
  );
}

export default DeleteConfirmationModal;
