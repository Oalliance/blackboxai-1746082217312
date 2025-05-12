import React, { useState } from 'react';

const Step2ProfileSetup = ({ onNext, onBack }) => {
  const [profile, setProfile] = useState({
    username: '',
    bio: '',
  });

  const handleChange = (e) => {
    setProfile({
      ...profile,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onNext(profile);
  };

  return (
    <form onSubmit={handleSubmit} className="p-6 bg-white rounded-lg shadow-md max-w-md mx-auto">
      <h2 className="text-xl font-bold mb-4">Step 2: Profile Setup</h2>
      <label className="block mb-2 font-semibold" htmlFor="username">Username</label>
      <input
        id="username"
        name="username"
        type="text"
        className="w-full p-2 mb-4 border border-gray-300 rounded"
        value={profile.username}
        onChange={handleChange}
        required
      />
      <label className="block mb-2 font-semibold" htmlFor="bio">Bio</label>
      <textarea
        id="bio"
        name="bio"
        className="w-full p-2 mb-6 border border-gray-300 rounded"
        value={profile.bio}
        onChange={handleChange}
        rows="4"
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
          Next
        </button>
      </div>
    </form>
  );
};

export default Step2ProfileSetup;
