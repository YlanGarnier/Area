import { StyleSheet } from 'react-native';

const styles= StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#D0CBED',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100%',
        padding: 10,
    },
    title: {
        fontSize: 24,
        fontWeight: 'bold',
        marginBottom: 20,
        marginTop: "15%",
        color: "#7D72C0",
    },
    scroll: {
        width: '100%',
        paddingHorizontal: '10%',
        margin: 10,
    },
    linksContainer: {
        flexDirection: 'row',
        flexWrap: 'wrap',
        justifyContent: 'space-between',
        width: '100%',
    },
    links: {
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        width: '48%',
        height: 100,
        borderColor: '#CF93D9',
        borderWidth: 2,
        borderRadius: 15,
        backgroundColor: 'white',
        margin: 3,
        padding: 10,
    },
    linkText: {
        color: '#7D72C0',
        fontWeight: "bold",
        fontSize: 16,
        textAlign: 'center',
        marginTop: 10,
    },
    finish: {
        height: 50,
        width: "40%",
        padding: 10,
        backgroundColor: '#725A8A',
        borderRadius: 5,
        marginRight: 10,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 50,
    },
    finishText: {
        color: 'white',
        fontSize: 16,
        textAlign: 'center',
    },
});

export default styles;
