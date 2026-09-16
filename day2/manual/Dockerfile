# copying the python and the OS Images
FROM python:3.12-alpine

# Creating a working directory for our application.
# After this our control will get into this app folder
WORKDIR /app

# Copying all the things from the current local directory to 
# the container current directory that is app directory
COPY . .

# Installing all the dependencies in the container
RUN pip install -r requirements.txt

# Starting the application inside the container.
# This will run only with command docker run
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]

# Expose the port that the application listens on.
EXPOSE 8000