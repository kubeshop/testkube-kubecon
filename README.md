# Testkube Kubecon Resources

Resources for KubeCon NA 2022 - for free Testkube Plushies giveaway! 

# Free Plushies

Our team prepared some freebies to give away between stickers and T-shirts you can also find out our great Kubepeople Plushies
just like this one: 

![Testkube-Kubeperson](https://user-images.githubusercontent.com/30776/196876546-3e1a095e-c7a9-42f5-9f15-f0995649af11.png)


We have around 50 plushies to give away - so there is one rule: first come first served!


# How can I get it? 

It's not so simple to get our Plushie, you need to complete the challenge first! 

## The challenge! 

1. You'll need to create the Testkube based test first.
2. You can do it on your local Testkube cluster 
   or on our demo server using UI
3. You'll need to run test and pass your email address as EMAIL variable
4. Test will call our service with requested email and our Team in the booth will be able to check if challenge was completed (you'll be asked for passed email address).


Simple ;) 

## (Option 1) How to add test using our demo cluster? 



1. Goto Testkube demo: 

https://demo.testkube.io

2. Add Test 

<img width="1546" alt="Zrzut ekranu 2022-10-19 o 13 10 58" src="https://user-images.githubusercontent.com/30776/196675802-c761e935-4bf3-4cf9-9a72-452b8c2df4b9.png">

You'll need to pass our test details which is in our git repo: 
- your test name
- `type`: just choose postman/collection
- `source`: choose `Git file` as we've stored our test as exported Postman Collection


3. Configure details and pass your contact details in form of email address as Test Variable

<img width="616" alt="Zrzut ekranu 2022-10-19 o 12 22 20" src="https://user-images.githubusercontent.com/30776/196675893-c578e2e4-7837-416a-8bd6-099148180154.png">

- Don't forget to hit the "Save" button

4. Run the test
5. Visit us in our booth and check if there are still plushies left!

<img width="1559" alt="image" src="https://user-images.githubusercontent.com/30776/196677715-17bec406-74a2-4a8a-93ab-e2ef616e6f82.png">

## (Option 2) Run it in your Testkube cluster

1. Follow Testkube installations instructions at: https://kubeshop.github.io/testkube/installing
3. Create test with Testkube CLI
```sh
testkube create test --name MY_TEST_NAME --git-uri https://github.com/kubeshop/testkube-NAKubecon22-FreePlushie.git --git-path test.json --type postman/collection --test-content-type git-file --variable EMAIL="MY_EMAIL_ADDRESS"
``` 
replace `MY_TEST_NAME` and `MY_EMAIL_ADDRESS` with yours.

3. Run test with Testkube CLI
```sh
testkube run test -f MY_TEST_NAME
``` 
4. Visit us in our booth!



