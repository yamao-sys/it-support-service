import { createContext, FC, useCallback, useContext, useState } from "react";
import { PostSupporterSignUpRequest } from "@/apis";

type SupporterSignUpSetContextType = {
  updateSignUpInput: (params: Partial<PostSupporterSignUpRequest>) => void;
  clearIdentificationKey: (keyToRemove: "frontIdentification" | "backIdentification") => void;
};

type SupporterSignUpContextType = {
  supporterSignUpInputs: PostSupporterSignUpRequest;
};

export const SupporterSignUpContext = createContext<SupporterSignUpContextType>(
  {} as SupporterSignUpContextType,
);

export const SupporterSignUpSetContext = createContext<SupporterSignUpSetContextType>(
  {} as SupporterSignUpSetContextType,
);

export const useSupporterSignUpContext = () =>
  useContext<SupporterSignUpContextType>(SupporterSignUpContext);

export const useSupporterSignUpSetContext = () =>
  useContext<SupporterSignUpSetContextType>(SupporterSignUpSetContext);

export const SupporterSignUpProvider: FC<{ children: React.ReactNode }> = ({ children }) => {
  const [supporterSignUpInputs, setSupporterSignUpInputs] = useState<PostSupporterSignUpRequest>({
    firstName: "",
    lastName: "",
    email: "",
    password: "",
  });

  const updateSignUpInput = useCallback((params: Partial<PostSupporterSignUpRequest>) => {
    setSupporterSignUpInputs((prev: PostSupporterSignUpRequest) => ({ ...prev, ...params }));
  }, []);

  const clearIdentificationKey = (keyToRemove: "frontIdentification" | "backIdentification") => {
    setSupporterSignUpInputs((prev: PostSupporterSignUpRequest) => {
      const { [keyToRemove]: _, ...rest } = prev; // eslint-disable-line @typescript-eslint/no-unused-vars
      return rest;
    });
  };

  return (
    <SupporterSignUpContext.Provider value={{ supporterSignUpInputs }}>
      <SupporterSignUpSetContext.Provider value={{ updateSignUpInput, clearIdentificationKey }}>
        {children}
      </SupporterSignUpSetContext.Provider>
    </SupporterSignUpContext.Provider>
  );
};
