import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
    modalOverlay: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'rgba(0,0,0,0.5)',
    },
    dropdownModal: {
        width: '85%',
        padding: 20,
        backgroundColor: '#FFFFFF',
        borderRadius: 15,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.2,
        shadowRadius: 10,
        elevation: 3,
    },
    iconRow: {
        flexDirection: 'row',
        justifyContent: 'space-around',
        marginVertical: 10,
    },
    iconInRow: {
        flex: 1,
        alignItems: 'center',
    }
});

export default styles;
