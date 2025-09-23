![Flash](https://icbotbz40ngrmv6r.public.blob.vercel-storage.com/flash.png)

# ⚡ Flash - Modern Internet Speed Testing Tool

Flash is a beautiful, modern web application that provides accurate internet speed testing with a sleek dark-themed interface. Built with Go and Fiber framework, featuring real-time speed measurements, responsive design, and smooth animations for an exceptional user experience.

## ✨ Features

### 🚀 **Speed Testing**
- **Accurate measurements**: Precise download and upload speed testing
- **Real-time progress**: Beautiful animated progress bars with gradient effects
- **Public IP detection**: Displays your current public IP address
- **Instant results**: Fast, reliable speed calculations with proper unit conversion

### 🎨 **Modern UI/UX**
- **Dark theme**: Elegant dark interface with gradient backgrounds
- **Responsive design**: Optimized for desktop, tablet, and mobile devices
- **Horizontal layout**: Professional dashboard-style layout for better data visualization
- **Smooth animations**: Fluid transitions and progress animations
- **Glass morphism**: Modern translucent card designs with backdrop blur

### 🛠 **Technical Stack**
- **Backend**: Go with Fiber framework for high performance
- **Frontend**: Vanilla HTML5, CSS3, and JavaScript
- **Styling**: Modern CSS with custom animations and gradients
- **Typography**: Inter font family for clean, professional appearance
- **Icons**: Beautiful SVG icons for enhanced visual appeal

## 🖥️ **Interface**

The application features a clean, horizontal layout with three main sections:
- **Your Public IP**: Shows your current public IP address with network icon
- **Download Speed**: Real-time download speed testing with progress visualization
- **Upload Speed**: Upload speed measurement with animated progress bars

## 🚀 **Getting Started**

### Prerequisites
- Go 1.19 or higher
- Modern web browser

### Installation

1. **Clone the repository:**
```bash
git clone https://github.com/harmoniousmoss/flash.git
cd flash
```

2. **Install dependencies:**
```bash
go mod tidy
```

3. **Run the application:**
```bash
go run main.go
```

4. **Open your browser:**
Navigate to [http://localhost:8080](http://localhost:8080)

### Development Mode

For development with auto-reload using Air:
```bash
air
```

## 📁 **Project Structure**

```
flash/
├── main.go                    # Main application entry point
├── handlers/                  # API handlers
│   ├── speedtest_handler.go  # Speed test logic
│   ├── ip_handler.go         # Public IP detection
│   ├── server_handler.go     # Server management
│   ├── test_runner.go        # Test execution engine
│   └── types.go              # Data structures
├── views/
│   └── index.html            # Main application interface
├── static/css/
│   └── styles.css            # Modern styling and animations
├── go.mod                    # Go module dependencies
└── README.md                 # This file
```

## 🎯 **API Endpoints**

- `GET /` - Main application interface
- `GET /api/speedtest` - Perform speed test and return results
- `GET /css/styles.css` - Application styling

## 🌟 **Key Improvements**

### Performance
- Optimized speed test algorithms for accurate measurements
- Efficient Go backend with minimal resource usage
- Fast loading times with optimized assets

### User Experience
- Intuitive interface with clear visual hierarchy
- Responsive design that works on all devices
- Smooth animations that enhance user engagement
- Professional appearance suitable for business use

### Technical Excellence
- Clean, maintainable code structure
- Modern web standards compliance
- Secure implementation with proper error handling
- Scalable architecture for future enhancements

## 🎨 **Design Philosophy**

Flash combines functionality with aesthetics, featuring:
- **Minimalist approach**: Clean interface without unnecessary clutter
- **Modern aesthetics**: Dark theme with subtle gradients and glass effects
- **Professional appearance**: Suitable for both personal and business use
- **Accessibility**: High contrast ratios and readable typography
- **Performance**: Smooth 60fps animations and fast load times

## 📱 **Responsive Design**

The application adapts seamlessly across devices:
- **Desktop**: Full horizontal layout with three-column grid
- **Tablet**: Responsive grid that adjusts to screen size
- **Mobile**: Optimized vertical layout for touch interaction

## 🛡️ **Browser Support**

Flash works on all modern browsers:
- Chrome 80+
- Firefox 75+
- Safari 13+
- Edge 80+

## 🤝 **Contributing**

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and feature requests.

## 📄 **License**

This project is open source and available under standard open source terms.

---

**Built with ❤️ by Syaeful Bahri**