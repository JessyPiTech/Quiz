# Makefile

# Variables
SRC_DIR := src
DOC_DIR := doc

# Nom du fichier exécutable
EXECUTABLE := Quiz.exe

# Compiler avec les options de build
BUILD_FLAGS :=

# Commandes
GO := go
MKDIR_P := mkdir -p
RM := rm -rf
MV := mv

# Cibles par défaut
all: build

# Cible pour construire le binaire et déplacer les dossiers
build: move_assets $(EXECUTABLE) 

# Règle pour déplacer les dossiers
move_assets:
	@$(MV) $(SRC_DIR)/players .
	@$(MV) $(SRC_DIR)/asset .
	


# Règles de construction
$(EXECUTABLE): $(wildcard $(SRC_DIR)/*.go)
	$(GO) build $(BUILD_FLAGS) -o $@ $(wildcard $(SRC_DIR)/*.go)

# Cible pour générer la documentation
doc:
	$(MKDIR_P) $(DOC_DIR)
	$(GO) doc -all -html -o $(DOC_DIR)/asset/index.html $(SRC_DIR)


# Cible pour nettoyer les fichiers générés et les assets
clean: clean_assets
	$(RM) $(EXECUTABLE)

# Règle pour remettre les dossiers à leur place
clean_assets:
	@$(MV) players $(SRC_DIR)/
	@$(MV) asset $(SRC_DIR)/


# Cible pour exécuter l'application
run: 
# Cible pour exécuter l'application
run: build
ifeq ($(filter-out run,$(MAKECMDGOALS)),)  # Vérifie s'il y a des arguments supplémentaires
	@echo "Tu n'as pas besoin de fichier texte en parametre, j'ai integre le choix au jeu."
else
	./$(EXECUTABLE) $(filter-out $@,$(MAKECMDGOALS))
endif
# Cible pour afficher l'aide
help:
	@echo "Usage:"
	@echo "  make          : Construire le binaire"
	@echo "  make doc      : Générer la documentation"
	@echo "  make clean    : Nettoyer les fichiers générés"
	@echo "  make run      : Construire et exécuter l'application"
	@echo "  make help     : Afficher cette aide"