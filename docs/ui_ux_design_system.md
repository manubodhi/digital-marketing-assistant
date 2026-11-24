# DOC-004: UI/UX Design System & Theme Specification

```yaml
id: DOC-004
type: Design-System-Specification
status: Accepted
domain: Frontend Presentation, Theming, Component Tokens
revision: 1.0
last_updated: 2026-11-24
dependencies: [DOC-001, DOC-002]
```

---

## 1. Visual Identity & Aesthetic Direction
- **Style:** Clean, modern, high-density SaaS management console (inspired by Linear and Vercel).
- **Core Goal:** Minimize cognitive clutter so marketers can quickly preview creatives, review schedules, and monitor ad spend.
- **Theme Modes:** First-class Dark and Light modes using CSS variables and semantic Tailwind color tokens.

---

## 2. Color Palette & Semantic Tokens

All UI widgets must use semantic CSS tokens instead of hardcoded hex values:

```css
/* Design Tokens (HSL format for dynamic alpha blending) */
:root {
  /* Light Mode */
  --background: 0 0% 100%;             /* Pure White */
  --foreground: 240 10% 3.9%;          /* Near Black */
  --card: 0 0% 100%;
  --card-foreground: 240 10% 3.9%;
  --popover: 0 0% 100%;
  --popover-foreground: 240 10% 3.9%;
  --primary: 240 5.9% 10%;             /* Deep Slate / Midnight */
  --primary-foreground: 0 0% 98%;
  --secondary: 240 4.8% 95.9%;         /* Soft Neutral */
  --secondary-foreground: 240 5.9% 10%;
  --muted: 240 4.8% 95.9%;
  --muted-foreground: 240 3.8% 46.1%;
  --accent: 240 4.8% 95.9%;
  --accent-foreground: 240 5.9% 10%;
  --destructive: 0 84.2% 60.2%;        /* Warning / Error Red */
  --destructive-foreground: 0 0% 98%;
  --border: 240 5.9% 90%;
  --input: 240 5.9% 90%;
  --ring: 240 5.9% 10%;
  --radius: 0.5rem;                    /* 8px default border radius */
}

.dark {
  /* Dark Mode (High-contrast slate/zinc) */
  --background: 240 10% 3.9%;          /* Deep Charcoal #09090b */
  --foreground: 0 0% 98%;
  --card: 240 10% 5.5%;                /* Elevated card background */
  --card-foreground: 0 0% 98%;
  --popover: 240 10% 5.5%;
  --popover-foreground: 0 0% 98%;
  --primary: 0 0% 98%;                 /* Crisp White action buttons */
  --primary-foreground: 240 5.9% 10%;
  --secondary: 240 3.7% 15.9%;
  --secondary-foreground: 0 0% 98%;
  --muted: 240 3.7% 15.9%;
  --muted-foreground: 240 5% 64.9%;
  --accent: 240 3.7% 15.9%;
  --accent-foreground: 0 0% 98%;
  --destructive: 0 62.8% 30.6%;
  --destructive-foreground: 0 0% 98%;
  --border: 240 3.7% 15.9%;
  --input: 240 3.7% 15.9%;
  --ring: 240 4.9% 83.9%;
}
```

---

## 3. Typography Scale

- **Primary Font Family:** Inter or Geist Sans (`font-sans`)
- **Monospace Font (for tokens & API keys):** JetBrains Mono or Geist Mono (`font-mono`)

| Role | Size | Weight | Line Height | Tailwind Class |
| :--- | :--- | :--- | :--- | :--- |
| **Page Title (H1)** | 24px (1.5rem) | 600 (Semibold) | 2rem | `text-2xl font-semibold tracking-tight` |
| **Section Header (H2)**| 18px (1.125rem) | 600 (Semibold) | 1.75rem | `text-lg font-semibold` |
| **Widget Label (H3)** | 14px (0.875rem) | 500 (Medium) | 1.25rem | `text-sm font-medium` |
| **Body Text** | 14px (0.875rem) | 400 (Regular) | 1.25rem | `text-sm text-foreground` |
| **Caption / Subtext** | 12px (0.75rem) | 400 (Regular) | 1rem | `text-xs text-muted-foreground` |

---

## 4. Component Layout Grid & Spacing Standard

1. **Dashboard Shell Layout:**
   - Left Sidebar (Fixed 240px width, collapsible to 64px icon rail).
   - Top Bar (Height 56px, contains Workspace Switcher, Notifications, User Menu, Theme Toggle).
   - Main Content Area (Fluid width, max-w-7xl centered with `px-8 py-6`).
2. **Spacing Grid:** Multiples of 4px (`p-1` = 4px, `p-2` = 8px, `p-4` = 16px, `p-6` = 24px).
3. **Widget State Mandate:**
   - Every widget MUST implement:
     - **Default State**: Clean card container (`border border-border bg-card rounded-lg p-6`).
     - **Loading State**: Shimmer skeleton placeholder (`animate-pulse bg-muted rounded`).
     - **Empty State**: Friendly icon, explanatory header, and a single primary action button.
     - **Error State**: Non-disruptive banner or input border (`border-destructive`).
