# Exam GO

* ### Fonction **IsAlpha(r rune)bool** <br>
    * **Description** <br> La fonction a pour but de retourner si oui ou non une rune 'r' est une lettre.

    * **Fonctionnement** <br> La fonction regarde et compare la valeur ASCII du la rune 'r'. <br> Si elle correspond à une valeur de lettre (de 65 à 90 ou 97 à 122) alors la fonction retourne ```true``` sinon elle retourne ```false```

    * **Example** <br> 
    ```go 
    fmt.Println(exam.IsAlpha('a'))
	fmt.Println(exam.IsAlpha('R'))
	fmt.Println(exam.IsAlpha('2'))
    ```
    ```go
    true  // 'a' est une lettre
    true  // 'R' est une lettre
    false // '2' n'est pas une lettre
    ```

* ### Fonction **Average(t []int)int**
    * **Description** <br> La fonction à pour but de calculer la moyenne des valeurs contenues dans un tableau de nombre 't'

    * **Fonctionnement** <br> La fonction va boucler sur tous les nombres du tableau en les ajoutant les uns aux autres. <br> Elle va ensuite diviser le total par le nombre de fois ou elle à été modifiée. <br> Dans le cas où le tableau contient ```-1```, cette valeur ne sera pas prise en compte dans l'ajout ni dans la division.

    * **Example**
    ```go
    fmt.Println(exam.Average([]int{14, 12, 7, 5, 17}))
	fmt.Println(exam.Average([]int{13, -1, 8, 15, 15}))
    ``` 
    ```go
    11 // Ici, 14+12+7+5+17 = 55 (total); On divise le total par le nombre de nombres positifs contenue dans le tableau donc 55/5=11

    12 // Ici, 13+8+15+15 = 51 (total); On divise le total par le nombre de nombres positifs contenue dans le tableau donc 51/4=12
    ```

* ### Fonction **PrintElements(tab []int, str []string)string**
    * **Description** <br> La fonction retourne un string contenant toutes les valeurs du tableau 'str' correspondant aux indexs contenus dans 'tab'

    * **Fonctionnement** <br> La fonction vas boucler sur 'tab' et gardant la valeur dans 'value' <br> 
    Elle vas ensuite ajouter au string final la valeur de str à l'index 'value'<br>
    
    * **Example**
    ```go
    fmt.Println(exam.PrintElements([]int{1, 4}, []string{"Choice", "Hello", "Solid", "Curtain", "There", "Forward"}))
	fmt.Println(exam.PrintElements([]int{2, 0, 7}, []string{"Kenobi", "Unity", "General", "Therapist"}))
    ```
    ```go
    Hello There
    General Kenobi !
    ``` 
    ![alt text](https://i.kym-cdn.com/photos/images/newsfeed/001/475/420/c62.gif)

* ### Programme **use-arguments**
    * **Description** <br> Ce programme a pour but d'afficher les arguments de lancement utilisés dans le format: <br> ```*Argument %d : %s\n* ``` ou %d est l'index de l'argument et %s est la valeur de l'argument

    * **Fonctionnement** <br>
    Le programme vas boucler sur les arguments de lancement ```os.Args[1:]``` <br> et vas ensuite les affichés dans la console avec le format voulu.

    * **Example** <br>
    ```go
     go run . "zdani" "daas7" "fae8adcsd3a" "yup" "zkr ni"
    ```
    ```go
    Argument 1 : zdani
    Argument 4 : yup
    Argument 5 : zkr ni

    ```

* ### Fonction **RecursiveSequence(current, max int)int**
    * **Description** <br> La fonction à pour but de moduler un nombre 'current' d'une façon differente si il est paire ou non tant que le prochain calcul ne dépasse pas 'max' sinon elle le retourne. <br> Régles de calcul: <br>
        * Si paire: current + 5
        * Si inpaire: current * 3 + 1

    * **Fonctionnement** <br> La fonction verifie si 'current' est plus petit que 'max' sinon elle retourne 0. <br> Elle vas ensuite calculer le 'next' en fonction de si le nombre est paire ou non. <br> Si le 'next' est plus grand que 'max' alors elle retourne 'current' <br>
    Sinon elle retourne le résultat donner en ce rappellant tel que: <br>
    ```return RecursiveSequence(next, max)```

    * **Example** <br>
    ```go
    fmt.Println(exam.RecursiveSequence(3, 54))
	fmt.Println(exam.RecursiveSequence(4, 73))
    ```
    ```go
    51
    33
    ```

* ### Fonction **Decipher(s string)string**
    * **Description** <br> La fonction permet de décrypter un string 's' crypter avec l'algorithme ROT16 (Change un caractère avec le 16ème qui le suit).

    * **Fonctionnement** <br> La fonction vas prendre en compte les valeurs ASCII de chaque caractères de 's' et lui enlever 16.<br>
    Elle vas ensuite verifier que le caractère calculer est bien une lettre sinon, cela veux dire que nous somme sortie des lettres voulu. <br> Si c'est le cas on vas donc partir de la fin de l'alphabet pour enlever la difference qu'il y as entre la valeus ASCII trouver et le début de l'alphabet 

    * **Example** <br>
    ```go
    fmt.Println(exam.Decipher("Zu ikyi kdu fuhieddu unsufjyeddubbu"))
    ```
    ```go
    Je suis une personne exceptionnelle
    ```
![alt text](https://c.tenor.com/jopMoI1TAgEAAAAC/steve-carell-the-office.gif)

* ### Fonction **Diamond(index int)**
    * **Description** <br> La fonction permet de dessiner un diamant dans le terminal avec une largeur maximal égale à 'index'

    * **Fonctionnement** <br> La fonction vas construire ligne par ligne le diamant en placant des '*' au extrémités et des espace sur le reste. <br> Elle vas ensuite afficher la ligne actuelle de la boucle

    * **Example** <br>
    ```go
    arg1 := 5
	exam.Diamond(arg1)
    ```
    ```go
        *
       * *
      *   *
     *     *
    *       *
     *     *
      *   *
       * *
        *
    ```
* ### Fonction **Conway(index int)int**
    * **Description** <br> La fonction vas retourner la valeurs de l'index 'index' de la suite du Conway <br>
    *https://fr.wikipedia.org/wiki/Suite_de_Conway*

    * **Fonctionnement** <br> La j'avou j'ai pas envie <br> *https://ytrack.learn.ynov.com/git/salexand/exam-go/src/branch/master/conway.go*

    * **Example** <br>
    ```go
    arg1 := 5
	fmt.Println(exam.Conway(arg1))
    ```
    ```go
    111221
    ```
*N'Oublie pas le like si t'as kiffé et abonnes-toi pour plus de Readme de qualité*

![alt text](https://thumbs.gfycat.com/DazzlingIckyAtlanticbluetang-max-1mb.gif)