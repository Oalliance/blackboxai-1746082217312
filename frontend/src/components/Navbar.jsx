import React from 'react';

const Navbar = () => {
  return (
    <nav className="bg-gray-800 p-4 text-white">
      <div className="container mx-auto flex justify-between items-center">
        <div className="text-lg font-bold">AppName</div>
        <div>
          <a href="/" className="px-3 hover:text-gray-300">Home</a>
          <a href="/dashboard" className="px-3 hover:text-gray-300">Dashboard</a>
          <a href="/profile" className="px-3 hover:text-gray-300">Profile</a>
          <a href="/login" className="px-3 hover:text-gray-300">Login</a>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
