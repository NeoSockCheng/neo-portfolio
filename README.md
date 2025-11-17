# Portfolio Website

A modern, responsive portfolio website built with Go and Tailwind CSS.

## Features

- Clean and modern design
- Responsive layout (mobile-friendly)
- Fast and lightweight
- Easy to customize
- Ready for Vercel deployment

## Tech Stack

- **Backend**: Go with Chi router
- **Frontend**: HTML templates with Tailwind CSS
- **Deployment**: Vercel
- **Styling**: Tailwind CSS (CDN)

## Project Structure

```
portfolio/
├── cmd/web/              # Application entry point
├── internal/             # Internal packages
│   ├── handlers/         # HTTP handlers
│   ├── models/           # Data models
│   ├── middleware/       # Custom middleware
│   └── config/           # Configuration
├── web/                  # Frontend assets
│   ├── templates/        # HTML templates
│   └── static/           # CSS, JS, images
├── data/                 # JSON data files
└── vercel.json          # Vercel configuration
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd portfolio
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
make run
# or
go run cmd/web/main.go
```

4. Open your browser and navigate to:
```
http://localhost:8080
```

## Customization

### Update Your Information

1. **Home Page**: Edit `web/templates/pages/home.html`
   - Update name, description, and intro text

2. **About Page**: Edit `web/templates/pages/about.html`
   - Add your bio, skills, and experience

3. **Projects**: Edit `data/projects.json`
   - Add your projects with descriptions, technologies, and links

4. **Contact Info**: Edit `web/templates/pages/contact.html` and `web/templates/partials/footer.html`
   - Update email, GitHub, and LinkedIn links

### Styling

- Main CSS: `web/static/css/styles.css`
- Tailwind is loaded via CDN in `web/templates/layouts/base.html`
- Customize colors and fonts in the base template

## Deployment to Vercel

1. Push your code to GitHub

2. Import your repository in Vercel:
   - Go to [vercel.com](https://vercel.com)
   - Click "New Project"
   - Import your GitHub repository

3. Vercel will auto-detect the Go configuration

4. Deploy!

Your site will be live at `https://your-project.vercel.app`

## Available Commands

```bash
make run      # Run the application
make build    # Build the application binary
make clean    # Clean build artifacts
make dev      # Run with auto-reload (requires Air)
```

## Environment Variables

Copy `.env.example` to `.env` and update:

```bash
cp .env.example .env
```

## Adding Features

### Email Contact Form

To enable the contact form to send emails:

1. Install an email package (e.g., gomail)
2. Update `internal/handlers/contact.go`
3. Add SMTP credentials to `.env`

### Database Integration

To add a database:

1. Add database driver to `go.mod`
2. Update `internal/models/` with database models
3. Add database configuration to `internal/config/`

## License

MIT License - feel free to use this for your own portfolio!

## Support

If you have questions or need help, please open an issue.
