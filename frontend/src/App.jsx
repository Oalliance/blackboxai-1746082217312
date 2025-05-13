import React, { useState } from 'react';
import Navigation from './components/Navigation';

const App = () => {
  const [currentPage, setCurrentPage] = useState('dashboard');

  const renderPage = () => {
    switch (currentPage) {
      case 'dashboard':
        return <div className="p-6">Welcome to the Dashboard</div>;
      case 'membership':
        return <div className="p-6">Membership Management</div>;
      case 'marketplace':
        return <div className="p-6">Marketplace Listings</div>;
      case 'governance':
        return <div className="p-6">Governance Proposals</div>;
      case 'disputes':
        return <div className="p-6">Dispute Resolution</div>;
      case 'profile':
        return <div className="p-6">User Profile</div>;
      default:
        return <div className="p-6">Page Not Found</div>;
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navigation currentPage={currentPage} onNavigate={setCurrentPage} />
      <main>{renderPage()}</main>
    </div>
  );
};

export default App;
