# My Susu  

**My Susu** is a mobile and USSD-based savings and financial management platform designed for market women and small business owners in Liberia and West Africa.  

## Overview  
My Susu empowers users by providing tools to:  
- Track their daily income and expenses.  
- Save a portion of their earnings for reinvestment or personal growth.  
- Manage finances effectively using mobile money platforms like MTN Mobile Money.  

The app is tailored for users with smartphones (mobile app) and non-smartphones (USSD platform), ensuring accessibility for a wide range of users.  

---

## Features  

### Mobile App  
- User authentication (Email and Password).  
- Dashboard with an overview of income, expenses, and savings balance.  
- Log daily income and expenses.  
- Manage savings (deposit and withdraw).  
- Push notifications for transaction updates.  

### USSD System  
- Log income and expenses via USSD.  
- Check savings balance.  
- Withdraw savings using mobile money.  

---

## Technologies Used  

### Mobile App  
- Framework: React Native  
- Notifications: Expo Push Notifications  

### Backend  
- Language: Golang  
- Router: Chi  
- Database Management: PostgreSQL, SQLC, and Goose  

### Integrations  
- Mobile Money: MTN Mobile Money API (Sandbox environment)  

---

## Development Plan  

### Week 1:  
- Backend setup with Golang, Chi, SQLC, and Goose.  
- Design database schema and implement migrations.  
- Develop and test core backend APIs (authentication, transactions, savings).  

### Week 2:  
- Build the mobile app using React Native.  
- Implement user flows (login, dashboard, savings, transactions).  
- Develop the USSD backend flow.  
- Integrate MTN Mobile Money API.  

---

## Challenges Identified  
- Ensuring seamless integration with MTN Mobile Money.  
- Balancing features between USSD and the mobile app.  
- Keeping the interface simple yet effective for users with limited technical knowledge.  

---

## Future Enhancements  
- Expand mobile money integrations (e.g., Orange Money, Airtel).  
- Add group savings (community susu).  
- Advanced financial analytics and tips.  

---

## Contributing  
For now, this is a solo project by **Cindy Tetama Johnson**. Contributions are welcome after the MVP launch.  

## License  
This project is licensed under the [MIT License](LICENSE).  
