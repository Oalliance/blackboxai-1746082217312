import React from 'react';

const Step4Confirmation = ({ onBack, onSubmit }) => {
  return (
    <div className="p-6 bg-white rounded-lg shadow-md max-w-md mx-auto">
      <h2 className="text-xl font-bold mb-4">Step 4: Confirmation</h2>
      <p className="mb-6">Please review your information and confirm to complete the onboarding process.</p>
      <div className="flex justify-between">
        <button
          onClick={onBack}
          className="bg-gray-300 text-gray-700 py-2 px-4 rounded hover:bg-gray-400 transition"
        >
          Back
        </button>
        <button
          onClick={onSubmit}
          className="bg-green-600 text-white py-2 px-4 rounded hover:bg-green-700 transition"
        >
          Confirm
        </button>
      </div>
    </div>
  );
};

export default Step4Confirmation;
