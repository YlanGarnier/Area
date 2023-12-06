import * as Linking from "expo-linking";
import * as SecureStore from "expo-secure-store";

function handleOpenURL async (url:any) => {
    const { hostname, path, queryParams} = Linking.parse(url);
    if (path === 'Lobby') {
        navigation.navigate("Lobby");
        await SecureStore.setItemAsync('DiscordToken', queryParams.token);
    }
}

export default handleOpenURL;
