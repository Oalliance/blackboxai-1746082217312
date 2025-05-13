import React from 'react';

const Navigation = ({ currentPage, onNavigate }) => {
  const navItems = [
    { id: 'dashboard', label: 'Dashboard' },
    { id: 'membership', label: 'Membership' },
    { id: 'marketplace', label: 'Marketplace' },
    { id: 'governance', label: 'Governance' },
    { id: 'disputes', label: 'Disputes' },
    { id: 'profile', label: 'Profile' },
  ];

  return (
    <nav className="bg-white shadow-md">
      <ul className="flex space-x-4 p-4">
        {navItems.map((item) => (
          <li key={item.id}>
            <button
              onClick={() => onNavigate(item.id)}
              className={`text-gray-700 hover:text-blue-600 font-medium ${
                currentPage === item.id ? 'border-b-2 border-blue-600' : ''
              }`}
              aria-current={currentPage === item.id ? 'page' : undefined}
            >
              {item.label}
            </button>
          </li>
        ))}
      </ul>
    </nav>
  );
};

export default Navigation;
