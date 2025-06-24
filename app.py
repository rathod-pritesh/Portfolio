from flask import Flask, url_for, render_template, redirect, request, flash, session, send_from_directory
from flask_sqlalchemy import SQLAlchemy
from datetime import datetime
import sqlite3
import os

app = Flask(__name__)
app.config['SECRET_KEY'] = os.urandom(24)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
app.config['SQLALCHEMY_TRACK_MODIFICATION'] = False
db = SQLAlchemy(app)

class PortfolioContent(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    hero_name = db.Column(db.String(100))
    hero_title = db.Column(db.String(255))
    about_title = db.Column(db.String(100))
    about_content = db.Column(db.Text)
    about_highlights = db.Column(db.String(255))
    contact_title = db.Column(db.String(100))
    contact_email = db.Column(db.String(100))
    contact_linkedin = db.Column(db.String(255))
    contact_github = db.Column(db.String(255))

class Contact(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), nullable=False)
    email = db.Column(db.String(100), nullable=False)
    subject = db.Column(db.String(200), nullable=False)
    message = db.Column(db.Text, nullable=False)

with app.app_context():
    db.create_all()

def validate_admin(username, password):
    conn = sqlite3.connect('admin.db')
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM admin WHERE username = ? AND password = ?" ,(username, password))
    admin = cursor.fetchone()
    conn.close()
    return admin

@app.route('/admin/login', methods = ['GET', 'POST'])
def admin_login():
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        admin = validate_admin(username, password)

        if admin:
            session['admin'] = True
            flash('Login successful!', 'success')
            return redirect(url_for('admin_dashboard'))
        else:
            flash('Invalid credentials.', 'danger')

    return render_template('admin_login.html')

@app.route('/admin/dashboard')
def admin_dashboard():
    if not session.get('admin'):
        return redirect(url_for('admin_login'))
    
    content = PortfolioContent.query.first()
    return render_template('admin_dashboard.html', content=content)

@app.route('/admin/update_content', methods=['POST'])
def update_content():
    if not session.get('admin'):
        return redirect(url_for('admin_login'))
    
    content = PortfolioContent.query.first()
    if not content:
        content = PortfolioContent()

    # Only update the fields that exist in the form
    form = request.form

    if form.get('hero_name'):
        content.hero_name = form.get('hero_name')
    if form.get('hero_title'):
        content.hero_title = form.get('hero_title')

    if form.get('about_title'):
        content.about_title = form.get('about_title')
    if form.get('about_content'):
        content.about_content = form.get('about_content')
    if form.get('about_highlights'):
        content.about_highlights = form.get('about_highlights')

    if form.get('contact_title'):
        content.contact_title = form.get('contact_title')
    if form.get('contact_email'):
        content.contact_email = form.get('contact_email')
    if form.get('contact_linkedin'):
        content.contact_linkedin = form.get('contact_linkedin')
    if form.get('contact_github'):
        content.contact_github = form.get('contact_github')

    # Handle resume file upload
    resume_file = request.files.get('resume')
    if resume_file and resume_file.filename != '':
        resume_path = os.path.join(app.root_path, 'static', 'resumes','resume.pdf')
        resume_file.save(resume_path)

    # Handle Profile image upload
    profile_image = request.files.get('profile_image')
    if profile_image and profile_image.filename != '':
        image_path = os.path.join(app.root_path, 'static', 'images', 'Profile.png')
        profile_image.save(image_path)

    db.session.add(content)
    db.session.commit()
    flash("Content updated successfully!", "success")
    return redirect(url_for('admin_dashboard'))

@app.route('/admin/logout')
def admin_logout():
    session.pop('admin', None)
    flash('Logged out successfully!', 'info')
    return redirect(url_for('admin_login'))

@app.route('/')
def home():
    content = PortfolioContent.query.first()
    return render_template('base.html', content = content)

@app.route('/download-resume')
def download_resume():
    resume_folder = os.path.join(app.root_path, 'static', 'resumes')
    return send_from_directory(directory= resume_folder, path='resume.pdf',as_attachment= True)
    

@app.route('/about')
def about():
    return redirect('/#about')
    
@app.route('/projects')
def projects():
    return redirect('/#projects')
    
@app.route('/certificates')
def certificates():
    return redirect('/#certificates')
    
@app.route('/contact')
def contact():
    return redirect('/#contact')
    
@app.route('/submit_contact', methods =['POST'])
def submit_contact():
    if request.method == 'POST':
        name = request.form['name']
        email = request.form['email']
        subject = request.form['subject']
        message = request.form['message']

        new_contact = Contact(name=name, email=email, subject=subject, message=message)
        db.session.add(new_contact)
        db.session.commit()

        flash('Your message has been sent successfully!','success')
        return redirect(url_for('contact'))

@app.route('/view_contacts')
def view_contacts():
    contacts = Contact.query.all()
    return render_template('view_contacts.html', contacts=contacts)

if __name__ == '__main__':
    app.run(debug=True)