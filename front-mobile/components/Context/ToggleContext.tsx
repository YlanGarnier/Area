import { createContext, useContext, useState, ReactNode } from 'react';

// Type for the toggle state
interface ServiceToggleContextType {
    toggles: { [key: string]: boolean };
    setToggle: (key: string, value: boolean) => void;
}

const ServiceToggleContext = createContext<ServiceToggleContextType | undefined>(undefined);


export const useServiceToggle = () => {
    const context = useContext(ServiceToggleContext);
    if (!context) {
        throw new Error("useServiceToggle doit être utilisé à l'intérieur de ServiceToggleProvider");
    }
    return context;
};

interface ServiceToggleProviderProps {
    children: ReactNode;
}

// The provider component
export const ServiceToggleProvider: React.FC<ServiceToggleProviderProps> = ({ children }) => {
    const [toggles, setToggles] = useState<{ [key: string]: boolean }>({});

    const setToggle = (key: string, value: boolean) => {
        setToggles(prev => ({ ...prev, [key]: value }));
    };

    return (
        <ServiceToggleContext.Provider value={{ toggles, setToggle }}>
            {children}
        </ServiceToggleContext.Provider>
    );
};
