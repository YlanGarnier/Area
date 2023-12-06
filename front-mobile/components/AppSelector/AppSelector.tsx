import React from 'react';
import { TouchableOpacity } from 'react-native';
import { AppSelectionModal } from './AppSelectionModal';
import styles from "./style";

export const AppSelector = ({ appImages, selectedApp, setSelectedApp }) => {
    const selectedAppComponent = appImages.find((img) => img.id === selectedApp)?.component;

    return (
        <>
            <TouchableOpacity style={styles.appSelector} onPress={() => setSelectedApp(true)}>
                {selectedAppComponent}
            </TouchableOpacity>
            <AppSelectionModal
                appImages={appImages}
                setSelectedApp={setSelectedApp}
            />
        </>
    );
};