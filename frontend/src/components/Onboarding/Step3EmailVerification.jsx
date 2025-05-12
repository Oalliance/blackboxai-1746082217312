import React, { useState } from 'react';

const Step3EmailVerification = ({ onNext, onBack }) => {
  const [code, setCode] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    // Here you would verify the code
    onNext();
  };

  return (
    <form onSubmit={handleSubmit} className="p-6 bg-white rounded-lg shadow-md max-w-md mx-auto">
      <h2 className="text-xl font-bold mb-4">Step 3: Email Verification</h2>
      <label className="block mb-2 font-semibold" htmlFor="code">Verification Code</label>
      <input
        id="code"
        type="text"
        className="w-full p-2 mb-6 border border-gray-300 rounded"
        value={code}
        onChange={(e) => setCode(e.target.value)}
        required
      />
      <div className="flex justify-between">
        <button
          type="button"
          onClick={onBack}
          className="bg-gray-300 text-gray-700 py-2 px-4 rounded hover:bg-gray-400 transition"
        >
          Back
        </button>
        <button
          type="submit"
          className="bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 transition"
        >
          Verify
        </button>
      </div>
    </form>
  );
};

export default Step3EmailVerification;
