# Utilisation de Node.js comme base
FROM node:19-alpine

# Créer un répertoire de travail
WORKDIR /app

# Installer Expo CLI globalement
RUN npm install -g expo-cli

# Copier le package.json et le package-lock.json s'ils sont présents
COPY package*.json ./

# Installer les dépendances
RUN npm install

RUN npm install --save-dev @types/react-native

# Copier les fichiers du projet
COPY . .

# Exécuter l'application
CMD ["expo", "start"]