# Image Replacement Guide

This guide shows you where to add your images to replace the placeholders in your portfolio website.

## Image Locations

All images should be placed in the `web/static/images/` directory.

## Required Images

### 1. Profile Photo
- **File name**: `profile.jpg` or `profile.png`
- **Location**: `web/static/images/profile.jpg`
- **Recommended size**: 400x400 pixels (square)
- **Used on**: Home page hero section

To replace the placeholder:
1. Add your photo to `web/static/images/profile.jpg`
2. Edit `web/templates/pages/home.html` around line 36
3. Replace the gradient div with:
   ```html
   <img src="/static/images/profile.jpg" alt="Neo Sock Cheng" class="w-40 h-40 rounded-full shadow-xl object-cover">
   ```

### 2. Company Logos

Add these company logo files:

- **Grab**: `grab-logo.png` (64x64px recommended)
  - Location: `web/static/images/grab-logo.png`
  - Replace in: `web/templates/pages/about.html` line ~58

- **Ant International**: `ant-logo.png` (64x64px recommended)
  - Location: `web/static/images/ant-logo.png`
  - Replace in: `web/templates/pages/about.html` line ~84

- **Systum360**: `systum360-logo.png` (64x64px recommended)
  - Location: `web/static/images/systum360-logo.png`
  - Replace in: `web/templates/pages/about.html` line ~110

To replace each logo placeholder:
```html
<!-- Change from: -->
<div class="w-16 h-16 bg-green-100 dark:bg-green-900 rounded-lg flex items-center justify-center">
    <span class="text-2xl font-bold text-green-600 dark:text-green-400">G</span>
</div>

<!-- To: -->
<img src="/static/images/grab-logo.png" alt="Grab" class="w-16 h-16 object-contain rounded-lg">
```

### 3. Project Screenshots

Add these project screenshot files:

- **FideFund**: `fidefund.png` (800x450px recommended)
- **GreenStep**: `greenstep.png` (800x450px recommended)
- **Mock WorldFirst**: `mock-worldfirst.png` (800x450px recommended)
- **Spotify Detection**: `spotify-detection.png` (800x450px recommended)
- **Red Cliff**: `red-cliff.png` (800x450px recommended)

**Location**: All in `web/static/images/`

The project images are already configured in `data/projects.json` with the correct paths. Just add your images to the folder and they'll appear automatically!

## Quick Steps

1. **Collect your images**:
   - 1 profile photo
   - 3 company logos
   - 5 project screenshots

2. **Rename them** according to the names above

3. **Copy to** `web/static/images/` folder

4. **For profile photo and company logos**: Edit the HTML files to replace the placeholder divs with `<img>` tags

5. **For project screenshots**: They'll appear automatically once the images are in place!

## Image Tips

- Use PNG for logos (supports transparency)
- Use JPG for photos and screenshots (smaller file size)
- Optimize images before uploading (use TinyPNG or similar)
- Keep file sizes under 500KB each
- Use consistent aspect ratios for projects (16:9 works well)

## Testing

After adding images, refresh your browser at `http://localhost:8080` to see them!
