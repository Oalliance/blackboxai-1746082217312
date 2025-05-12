import React, { useState } from 'react';

const Step1UserDetails = ({ onNext }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    onNext({ name, email });
  };

  return (
    <form onSubmit={handleSubmit} className="p-6 bg-white rounded-lg shadow-md max-w-md mx-auto">
      <h2 className="text-xl font-bold mb-4">Step 1: User Details</h2>
      <label className="block mb-2 font-semibold" htmlFor="name">Name</label>
      <input
        id="name"
        type="text"
        className="w-full p-2 mb-4 border border-gray-300 rounded"
        value={name}
        onChange={(e) => setName(e.target.value)}
        required
      />
      <label className="block mb-2 font-semibold" htmlFor="email">Email</label>
      <input
        id="email"
        type="email"
        className="w-full p-2 mb-6 border border-gray-300 rounded"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        required
      />
      <button
        type="submit"
        className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
      >
        Next
      </button>
    </form>
  );
};

export default Step1UserDetails;
