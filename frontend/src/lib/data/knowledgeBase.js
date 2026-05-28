export const knowledgeBase = [
  // ─── Identity & Background ───────────────────────────────────────────────
  {
    tags: ["who", "about", "introduce", "yourself", "pritesh", "background"],
    question: "Who is Pritesh Rathod? / Tell me about yourself",
    answer:
      "Hey! I'm Pritesh Rathod — a backend-focused developer with a passion for building AI-powered applications, scalable APIs, and clean web systems. I'm currently pursuing my M.Sc. in Computer Applications & IT from Gujarat University. My go-to stack is Python, FastAPI, Golang, and SvelteKit, and I love exploring where AI meets real-world engineering.",
  },

  // ─── Skills ──────────────────────────────────────────────────────────────
  {
    tags: ["skills", "technologies", "tech stack", "know", "expertise", "proficient"],
    question: "What are your technical skills? / What technologies do you work with?",
    answer:
      "My technical toolkit spans a pretty wide range! On the backend I work with Python, Golang, FastAPI, and Flask. For frontends I use SvelteKit with TailwindCSS. In the AI/ML space I'm comfortable with LangChain, Pinecone, HuggingFace embeddings, Groq LLMs, OpenCV, and RAG pipelines. I also work with MySQL, MongoDB Atlas, AstraDB, Docker, and n8n for automation workflows.",
  },
  {
    tags: ["backend", "server", "api", "rest"],
    question: "What backend technologies do you use?",
    answer:
      "Backend is where I feel most at home! I primarily build with FastAPI and Flask in Python, and Golang for performance-critical services. I design RESTful APIs, work with MySQL and MongoDB Atlas for data persistence, and have explored AstraDB for NoSQL cloud use cases. I focus a lot on clean architecture and making systems that actually scale.",
  },
  {
    tags: ["frontend", "ui", "svelte", "tailwind", "web design"],
    question: "What frontend technologies do you work with?",
    answer:
      "For the frontend I mostly use SvelteKit paired with TailwindCSS — it's a combo I really enjoy because it keeps things fast and clean. I'm also comfortable with plain HTML/CSS, Bootstrap, JavaScript, and Vite for tooling. I care about building responsive, modern interfaces that complement the backend work.",
  },
  {
    tags: ["ai", "machine learning", "llm", "chatbot", "nlp", "rag", "langchain"],
    question: "What AI / ML technologies do you work with?",
    answer:
      "AI development is one of my favourite areas! I've built RAG (Retrieval-Augmented Generation) systems using LangChain and Pinecone, integrated HuggingFace embedding models, and used Groq LLMs for fast inference. I've also worked with OpenCV for computer vision tasks, built NLP-based voice assistants, and created AI chatbots that hook into real data sources.",
  },
  {
    tags: ["database", "mysql", "mongodb", "pinecone", "astradb", "nosql"],
    question: "What databases do you use?",
    answer:
      "I pick the database that fits the job best. For relational data I use MySQL, for flexible document storage it's MongoDB Atlas, for vector search in AI apps I use Pinecone, and I've also explored AstraDB for cloud-native NoSQL scenarios. I'm comfortable writing queries, designing schemas, and optimising for performance.",
  },
  {
    tags: ["frameworks", "fastapi", "flask", "sveltekit", "langchain"],
    question: "What frameworks do you use?",
    answer:
      "My most-used frameworks are FastAPI (my primary backend choice for its speed and automatic docs), Flask for simpler or prototype APIs, SvelteKit for full-stack web apps, LangChain for AI orchestration, and TailwindCSS for UI styling. Each project gets the right tool for its needs.",
  },

  // ─── Projects ─────────────────────────────────────────────────────────────
  {
    tags: ["projects", "built", "work", "portfolio", "showcase"],
    question: "What projects have you built?",
    answer:
      "I've built quite a few things I'm proud of! The highlights include: an AI Voice Assistant, a Medical Chatbot using RAG, BuildTrack (a construction management system), an AI Image Generator, my personal portfolio with an embedded AI chatbot, EcomTrack (an ecommerce analytics backend), and several automation workflows with n8n. Want details on any specific one?",
  },
  {
    tags: ["voice assistant", "ai voice", "speech", "nlp assistant"],
    question: "What is the AI Voice Assistant project?",
    answer:
      "The AI Voice Assistant is a full-stack web app I built with FastAPI on the backend and SvelteKit on the frontend. It supports real-time voice interaction, uses NLP for intent recognition, has a clean responsive UI, and includes user authentication. Basically, it lets you talk to an AI assistant right from your browser.",
  },
  {
    tags: ["medical chatbot", "rag", "pdf", "healthcare", "doctor bot"],
    question: "What is the Medical Chatbot project?",
    answer:
      "The Medical Chatbot is a RAG-based AI app that lets users upload medical PDF documents and then ask questions about them. Under the hood it uses LangChain to orchestrate the pipeline, Pinecone for vector storage, HuggingFace embeddings to encode the document content, Groq LLM for generating answers, and FastAPI to serve everything. It's a solid demonstration of document-grounded AI.",
  },
  {
    tags: ["buildtrack", "construction", "attendance", "face recognition", "management"],
    question: "What is BuildTrack?",
    answer:
      "BuildTrack is a construction management system I built using Flask, MySQL, OpenCV, and Bootstrap. It handles project tracking, generates reports, and — the coolest part — uses face recognition to automate attendance monitoring on-site. It was a great project for combining traditional web dev with computer vision.",
  },
  {
    tags: ["image generator", "text to image", "pollinations", "ai art"],
    question: "What is the AI Image Generator project?",
    answer:
      "The AI Image Generator is a fun app where you type a text prompt and it produces an AI-generated image. I built the backend with FastAPI and the frontend with SvelteKit, and it connects to Pollinations AI APIs to do the actual generation. Great for experimenting with creative prompts!",
  },
  {
    tags: ["portfolio", "personal site", "golang backend", "chatbot portfolio"],
    question: "Tell me about your personal portfolio project.",
    answer:
      "My portfolio is a full-stack project — SvelteKit on the frontend and Golang powering the backend REST APIs. It's connected to MongoDB Atlas for data, includes smooth UI animations, showcases all my projects, and has an AI-powered chatbot assistant built right in (that's me — the chatbot you're talking to!). It's essentially a living demonstration of my skills.",
  },
  {
    tags: ["ecomtrack", "ecommerce", "analytics", "astradb", "nosql"],
    question: "What is EcomTrack?",
    answer:
      "EcomTrack is a backend project I built to explore AstraDB — a cloud-native NoSQL database. It provides ecommerce analytics APIs and gave me hands-on experience with query-first data modelling and NoSQL architecture concepts. Think of it as a learning-driven project that turned into a solid backend system.",
  },
  {
    tags: ["automation", "n8n", "workflow", "google sheets", "news digest"],
    question: "What automation projects have you built?",
    answer:
      "I've built a couple of automation workflows using n8n. One automatically handles contact form submissions and pipes the data into Google Sheets. Another is a tech news digest system that scrapes, filters, and summarises relevant tech news automatically. These projects showed me how powerful no-code/low-code automation can be when combined with custom backend logic.",
  },

  // ─── Education ────────────────────────────────────────────────────────────
  {
    tags: ["education", "degree", "university", "study", "college", "qualification", "gujarat"],
    question: "What is your educational background?",
    answer:
      "I completed my B.Sc. in Computer Applications & IT from Gujarat University, and I'm currently pursuing my M.Sc. in Computer Applications & IT — also from Gujarat University. Academia and hands-on project work have gone hand-in-hand for me throughout my studies.",
  },

  // ─── Certifications ───────────────────────────────────────────────────────
  {
    tags: ["certifications", "courses", "certificates", "achievements", "anthropic", "python"],
    question: "What certifications do you have?",
    answer:
      "I've earned certifications in: AI Fluency: Framework & Foundations (by Anthropic), Advanced Python Programming, Go Programming Language, Figma UI/UX Design, and Java Mastery — Intermediate. I enjoy keeping my skills sharp through structured learning alongside real project work.",
  },

  // ─── Interests & Career ───────────────────────────────────────────────────
  {
    tags: ["job", "role", "career", "interest", "looking for", "hire", "open to", "opportunity"],
    question: "What kind of roles are you interested in?",
    answer:
      "I'm most interested in backend engineering roles, especially where AI integration, REST API design, or scalable system architecture is involved. Positions around AI-powered application development, automation engineering, or full-stack work with a strong backend focus are right in my wheelhouse. I love building things that are both technically solid and genuinely useful.",
  },
  {
    tags: ["hire", "available", "work with", "collaborate", "freelance"],
    question: "Are you open to new opportunities / can I hire you?",
    answer:
      "Absolutely, I'm open to discussing new opportunities! Whether it's a full-time role, freelance project, or collaboration, feel free to reach out. You can connect with me through GitHub or the contact form on this portfolio.",
  },

  // ─── GitHub / Links ───────────────────────────────────────────────────────
  {
    tags: ["github", "code", "source", "repository", "projects link"],
    question: "Where can I find your code / GitHub profile?",
    answer:
      "All my projects and source code are up on GitHub — check them out at: https://github.com/rathod-pritesh. Feel free to explore, star anything you find useful, or open an issue if you have questions!",
  },

  // ─── General / Catch-all ─────────────────────────────────────────────────
  {
    tags: ["contact", "reach", "email", "message", "get in touch"],
    question: "How can I contact Pritesh?",
    answer:
      "The best way to reach me is through the contact form on this portfolio, or you can find me on GitHub at https://github.com/rathod-pritesh. I'm always happy to chat about projects, opportunities, or just tech in general!",
  },
  {
    tags: ["golang", "go language", "go backend"],
    question: "Do you work with Golang?",
    answer:
      "Yes! Golang is one of my backend languages. I used it to build the REST API layer for my personal portfolio, connecting to MongoDB Atlas. I love Go for its performance, simplicity, and how well it handles concurrency — it's a great fit for building fast, reliable backend services.",
  },
  {
    tags: ["python", "why python", "python experience"],
    question: "Do you work with Python?",
    answer:
      "Python is my primary language and where I have the most depth. I use it for backend APIs (FastAPI, Flask), AI and ML pipelines (LangChain, HuggingFace, OpenCV), scripting, and automation. It's incredibly versatile and the ecosystem for AI development is unmatched.",
  },
  {
    tags: ["docker", "devops", "deployment", "containerisation"],
    question: "Do you have DevOps or Docker experience?",
    answer:
      "I work with Docker for containerising applications and managing consistent deployment environments. While my primary focus is backend and AI development, I understand the DevOps side enough to containerise services, write Dockerfiles, and manage basic deployment pipelines.",
  },
];