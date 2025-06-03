import { createContext, FC, useCallback, useContext, useState } from "react";
import { PostCompanySignUpRequest } from "@/apis";

type CompanySignUpSetContextType = {
  updateSignUpInput: (params: Partial<PostCompanySignUpRequest>) => void;
  clearIdentificationKey: (keyToRemove: "finalTaxReturn") => void;
};

type CompanySignUpContextType = {
  companySignUpInputs: PostCompanySignUpRequest;
};

export const CompanySignUpContext = createContext<CompanySignUpContextType>(
  {} as CompanySignUpContextType,
);

export const CompanySignUpSetContext = createContext<CompanySignUpSetContextType>(
  {} as CompanySignUpSetContextType,
);

export const useCompanySignUpContext = () =>
  useContext<CompanySignUpContextType>(CompanySignUpContext);

export const useCompanySignUpSetContext = () =>
  useContext<CompanySignUpSetContextType>(CompanySignUpSetContext);

export const CompanySignUpProvider: FC<{ children: React.ReactNode }> = ({ children }) => {
  const [companySignUpInputs, setCompanySignUpInputs] = useState<PostCompanySignUpRequest>({
    name: "",
    email: "",
    password: "",
  });

  const updateSignUpInput = useCallback((params: Partial<PostCompanySignUpRequest>) => {
    setCompanySignUpInputs((prev: PostCompanySignUpRequest) => ({ ...prev, ...params }));
  }, []);

  const clearIdentificationKey = (keyToRemove: "finalTaxReturn") => {
    setCompanySignUpInputs((prev: PostCompanySignUpRequest) => {
      const { [keyToRemove]: _, ...rest } = prev; // eslint-disable-line @typescript-eslint/no-unused-vars
      return rest;
    });
  };

  return (
    <CompanySignUpContext.Provider value={{ companySignUpInputs }}>
      <CompanySignUpSetContext.Provider value={{ updateSignUpInput, clearIdentificationKey }}>
        {children}
      </CompanySignUpSetContext.Provider>
    </CompanySignUpContext.Provider>
  );
};
