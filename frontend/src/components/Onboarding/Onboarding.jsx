import React, { useState } from 'react';
import Step1UserDetails from './Step1UserDetails';
import Step2ProfileSetup from './Step2ProfileSetup';
import Step3EmailVerification from './Step3EmailVerification';
import Step4Confirmation from './Step4Confirmation';

const Onboarding = () => {
  const [step, setStep] = useState(1);
  const [userDetails, setUserDetails] = useState({});
  const [profileSetup, setProfileSetup] = useState({});

  const nextStep = () => setStep(step + 1);
  const prevStep = () => setStep(step - 1);

  const handleUserDetailsNext = (data) => {
    setUserDetails(data);
    nextStep();
  };

  const handleProfileSetupNext = (data) => {
    setProfileSetup(data);
    nextStep();
  };

  const handleEmailVerificationNext = () => {
    nextStep();
  };

  const handleConfirmationSubmit = () => {
    // Submit all data to backend or perform final actions
    alert('Onboarding complete!');
  };

  switch (step) {
    case 1:
      return <Step1UserDetails onNext={handleUserDetailsNext} />;
    case 2:
      return <Step2ProfileSetup onNext={handleProfileSetupNext} onBack={prevStep} />;
    case 3:
      return <Step3EmailVerification onNext={handleEmailVerificationNext} onBack={prevStep} />;
    case 4:
      return <Step4Confirmation onBack={prevStep} onSubmit={handleConfirmationSubmit} />;
    default:
      return null;
  }
};

export default Onboarding;
