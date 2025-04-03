import { createContext, FC, useCallback, useContext, useState } from "react";
import { CompanySignUpInput } from "../_types";

type CompanySignUpSetContextType = {
  updateSignUpInput: (params: Partial<CompanySignUpInput>) => void;
  clearIdentificationKey: (keyToRemove: "finalTaxReturn") => void;
};

type CompanySignUpContextType = {
  companySignUpInputs: CompanySignUpInput;
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
  const [companySignUpInputs, setCompanySignUpInputs] = useState<CompanySignUpInput>({
    name: "",
    email: "",
    password: "",
  });

  const updateSignUpInput = useCallback((params: Partial<CompanySignUpInput>) => {
    setCompanySignUpInputs((prev: CompanySignUpInput) => ({ ...prev, ...params }));
  }, []);

  const clearIdentificationKey = (keyToRemove: "finalTaxReturn") => {
    setCompanySignUpInputs((prev: CompanySignUpInput) => {
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
