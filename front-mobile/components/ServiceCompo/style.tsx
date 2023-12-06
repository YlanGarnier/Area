import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
  container: {
    height: 70,
    justifyContent: "center",
    alignItems: "center",
    width: "95%",
    backgroundColor: "#f9f9f9",
    borderRadius: 15,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
    margin: 10,
  },
  content: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    width: "100%",
  },
  titleContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    width: "100%",
    marginBottom: 5,
  },
  leftContainer: {
    flexDirection: "row",
    position: "absolute",
    left: 30,
  },
  image: {
    width: 30,
    height: 30,
    marginRight: 5,
  },
  title: {
    fontSize: 20,
    color: "#333",
    maxWidth: "40%",
  },
  description: {
    top: 20,
    fontSize: 15,
    color: "#333",
    textAlign: "center",
    maxWidth: "80%",
  },
  deleteButton: {
    position: "absolute",
    top: 20,
    right: 20,
  },
});

export default styles;
