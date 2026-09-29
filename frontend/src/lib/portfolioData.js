export const defaultHome = {
  _id: "home",
  name: "Pritesh Rathod",
  role: "Python Backend Developer | AI & LLM Integrations | Golang Engineer",
  tagline: "Building scalable backend systems and AI-powered applications with FastAPI, LangChain, and LangGraph.",
  resumeUrl: "/resume/Pritesh_Rathod_resume.pdf",
  profileImage: "/images/profile_photo.png"
};

export const defaultAbout = {
  _id: "about",
  title: "About Me",
  description: "I am a Python Backend Developer and I am currently pursuing an M.Sc. in Computer Applications & Information Technology at Gujarat University. I have been working mainly with Python, FastAPI, and backend development, while also exploring AI technologies such as LangChain, LangGraph, RAG, and LLM-based applications."
};

export const defaultSkills = {
  _id: "skills",
  title: "Technical Skills",
  categories: [
    {
      name: "Backend Development",
      icon: "fa-solid fa-server",
      order: 1,
      skills: [
        { name: "Python", img: "/icons/python.svg" },
        { name: "FastAPI", img: "/icons/fastapi.svg" },
        { name: "Golang", img: "/icons/go.svg" },
        { name: "Gin", img: "/icons/gin.svg" }
      ]
    },
    {
      name: "AI & ML Orchestration",
      icon: "fa-solid fa-brain",
      order: 2,
      skills: [
        { name: "LangChain", img: "/icons/langchain.svg" },
        { name: "LangGraph", img: "/icons/langgraph.svg" },
        { name: "RAG", img: "/icons/rag.svg" }
      ]
    }
  ]
};

export const defaultProjects = [
  {
    title: "VideoRedact",
    description: "Automated video privacy application using YOLO and OpenCV to detect and blur faces and license plates.",
    image: "/images/VideoRedact.png",
    github: "https://github.com/rathod-pritesh",
    technologies: ["Python", "OpenCV", "YOLO", "FastAPI"],
    gradient: "from-blue-500 to-cyan-500",
    order: 1,
    isFeatured: true
  },
  {
    title: "Medical RAG Chatbot",
    description: "AI medical chatbot using RAG, Pinecone vector search, and Groq/LLaMA 3.",
    image: "/images/chatbot.png",
    github: "https://github.com/rathod-pritesh/end-to-end-medical-chatbot-generative-ai",
    technologies: ["Python", "LangChain", "Pinecone", "Groq"],
    gradient: "from-purple-500 to-indigo-500",
    order: 2,
    isFeatured: true
  },
  {
    title: "BuildTrack",
    description: "Construction management system with Flask, MySQL, and Bootstrap.",
    image: "/images/buildtrack.png",
    github: "https://github.com/rathod-pritesh/BuildTrack-construction-management-system",
    technologies: ["Flask", "MySQL", "Bootstrap"],
    gradient: "from-green-500 to-emerald-500",
    order: 3,
    isFeatured: true
  },
  {
    title: "Netflix Watchlist API",
    description: "High-performance RESTful API built with Go and MongoDB featuring modular MVC architecture.",
    image: "/images/netflix.png",
    github: "https://github.com/rathod-pritesh/go-mongodb-crud-api",
    technologies: ["Golang", "MongoDB"],
    gradient: "from-red-500 to-pink-500",
    order: 4,
    isFeatured: true
  }
];

export const defaultAutomations = [];

export const defaultCertifications = [
  {
    name: "AI Fluency: Framework & Foundations",
    company: "Anthropic",
    link: "https://priteshrathod.vercel.app/certificates/Anthropic.pdf",
    issueDate: "Mar 2026",
    order: 1
  },
  {
    name: "Google AI Professional Certificate",
    company: "Google",
    link: "https://priteshrathod.vercel.app/certificates/Google_AI.pdf",
    issueDate: "Jun 2026",
    order: 2
  }
];

export const defaultEducation = [
  {
    degree: "M.Sc Computer Applications & IT",
    institution: "K. S. School of Business Management and Information Technology",
    yearStart: "2025",
    yearEnd: "PRESENT",
    focus: "Focused on backend systems, AI integrations, automation workflows, and scalable application development.",
    order: 2
  },
  {
    degree: "B.Sc Computer Applications & IT",
    institution: "K. S. School of Business Management and Information Technology",
    yearStart: "2022",
    yearEnd: "2025",
    focus: "Built strong foundations in software engineering, databases, APIs, and modern web technologies.",
    order: 1
  }
];
